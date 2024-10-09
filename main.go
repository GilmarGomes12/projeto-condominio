package main

import (
    "condominio/config"
    "condominio/models"
    "condominio/routes"
    "database/sql"
    "log"
    "net/http"

    "go.uber.org/zap"
)

// Função principal que inicializa a aplicação.
func main() {
    // Inicializa a configuração (carrega variáveis de ambiente e inicializa o logger).
    if err := config.InitConfig(); err != nil {
        log.Fatalf("Erro na inicialização da configuração: %v", err)
    }

    // Conecta ao banco de dados.
    if err := config.ConnectDB(); err != nil {
        config.Logger().Fatal("Não foi possível conectar-se ao banco de dados", zap.Error(err))
    }
    defer config.CloseDB()

    // Cria as tabelas no banco de dados, se não existirem.
    createTables(config.GetDB())

    // Cria o roteador das rotas personalizadas.
    router := routes.Router(config.GetDB())

    // Middleware de logging.
    router.Use(loggingMiddleware)

    // Servir arquivos estáticos.
    router.PathPrefix("/templates/assets/").Handler(http.StripPrefix("/templates/assets/", http.FileServer(http.Dir("templates/assets"))))

    // Rota para criar usuário.
    router.HandleFunc("/create-user", createUserHandler(config.GetDB())).Methods("POST")

    // Inicia o servidor na porta 8080.
    config.Logger().Info("Iniciando servidor na porta :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        config.Logger().Fatal("Erro ao iniciar o servidor", zap.Error(err))
    }
}

// loggingMiddleware é um middleware que registra todas as requisições HTTP recebidas.
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        config.Logger().Info("Request", zap.String("method", r.Method), zap.String("uri", r.RequestURI))
        next.ServeHTTP(w, r)
    })
}

// createUserHandler é o handler para criar um novo usuário.
func createUserHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        username := r.FormValue("username")
        password := r.FormValue("password")

        usuario := &models.Usuario{
            Username: username,
            Password: password,
        }

        // Cria um novo usuário no banco de dados.
        err := usuario.Create(db)
        if err != nil {
            config.Logger().Error("Erro ao criar usuário", zap.Error(err))
            http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
            return
        }

        // Tenta associar a permissão ao usuário.
        err = models.AssignPermission(db, usuario.ID, "porteiro")
        if err != nil {
            config.Logger().Error("Erro ao associar permissão ao usuário", zap.Error(err))
            http.Error(w, "Erro ao associar permissão ao usuário", http.StatusInternalServerError)
            return
        }

        config.Logger().Info("Usuário criado com sucesso", zap.Int("userID", usuario.ID))
        http.Redirect(w, r, "/login", http.StatusSeeOther)
    }
}

// createTables cria as tabelas necessárias no banco de dados, se não existirem.
func createTables(db *sql.DB) {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS usuario (
            id SERIAL PRIMARY KEY,
            username VARCHAR(255) NOT NULL,
            password VARCHAR(255) NOT NULL
        );`,
        `CREATE TABLE IF NOT EXISTS permissoes (
            id SERIAL PRIMARY KEY,
            user_id INT NOT NULL,
            permissao VARCHAR(255) NOT NULL,
            FOREIGN KEY (user_id) REFERENCES usuario(id)
        );`,
    }

    // Executa cada query para criar as tabelas.
    for _, query := range queries {
        _, err := db.Exec(query)
        if err != nil {
            log.Fatalf("Erro ao criar tabela: %v", err)
        }
    }
}