package routes

import (
	"condominio/controllers"
	"condominio/middleware"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// Router configura todas as rotas da aplicação e retorna um roteador mux.Router.
func Router(db *sql.DB) *mux.Router {
    router := mux.NewRouter()

    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

    // Rota para criar um novo usuário.
    router.HandleFunc("/usuarios", controllers.CreateUserHandler(db))

    // Rota para a administração, requer permissão de 'admin'.
    router.Handle("/administracao", middleware.WithUserID(middleware.RequirePermission("admin", db, controllers.AdminHandler))).Methods("GET")

    // Rota para o síndico, requer permissão de 'sindico'.
    router.Handle("/sindico", middleware.WithUserID(middleware.RequirePermission("sindico", db, controllers.SindicoHandler))).Methods("GET")

    // Rota para o porteiro, requer permissão de 'porteiro'.
    router.Handle("/porteiro", middleware.WithUserID(middleware.RequirePermission("porteiro", db, controllers.PorteiroHandler))).Methods("GET")

    // Rota para a página inicial.
    router.HandleFunc("/index", controllers.IndexHandler).Methods("GET")

    // Rota para a página de login, permite métodos GET e POST.
    router.HandleFunc("/login", controllers.LoginHandler(db)).Methods("GET", "POST")

    // Rota para o menu principal, requer autenticação do usuário.
    router.HandleFunc("/menu", middleware.WithUserID(controllers.MenuHandler)).Methods("GET")

    // Rotas para o registro de usuário.
    router.HandleFunc("/register", controllers.RegisterFormHandler).Methods("GET")
    router.HandleFunc("/register", controllers.RegisterUserHandler(db)).Methods("POST")

    // Rotas para gerenciar moradores.
    router.HandleFunc("/morador", controllers.MoradorFormHandler).Methods("GET")
    router.HandleFunc("/morador", controllers.MoradorHandler(db)).Methods("POST")

    // Rotas para gerenciar agendamentos.
    router.HandleFunc("/agendamentos", controllers.AgendamentosFormHandler).Methods("GET")
    router.HandleFunc("/agendamentos", controllers.AgendamentosHandler(db)).Methods("POST")

    // Rotas para gerenciar visitantes.
    router.HandleFunc("/visitantes", controllers.VisitantesFormHandler).Methods("GET")
    router.HandleFunc("/visitantes", controllers.VisitantesHandler(db)).Methods("POST")

    // Rotas para gerenciar prestadores de serviço.
    router.HandleFunc("/prestadores", controllers.PrestadorFormHandler).Methods("GET")
    router.HandleFunc("/prestadores", controllers.PrestadorHandler(db)).Methods("POST")

    // Rotas para gerenciar encomendas.
    router.HandleFunc("/encomendas", controllers.EncomendasHandler(db)).Methods("POST", "GET")

    // Rotas para gerenciar ocorrências.
    router.HandleFunc("/ocorrencias", controllers.OcorrenciaHandler(db)).Methods("POST", "GET")

    // Rotas para gerenciar funcionários do condomínio.
    router.HandleFunc("/funcionarios", controllers.FuncionariosHandler(db)).Methods("POST", "GET")

    // Rota para gerenciar veículos.
    router.HandleFunc("/veiculos", controllers.VeiculosHandler())

    // Rotas para gerenciar funcionários domésticos.
    router.HandleFunc("/domesticos", controllers.DomesticosHandler(db)).Methods("POST", "GET")

    // Rotas para agendamento de mudanças.
    router.HandleFunc("/mudanca", controllers.MudancaHandler(db)).Methods("POST", "GET")

    // Rotas para configurações.
    router.HandleFunc("/configuracoes", controllers.ConfiguracoesHandler(db)).Methods("POST", "GET")

    // Rotas para pesquisa unificada.
    router.HandleFunc("/pesquisar", controllers.PesquisaHandler(db)).Methods("GET")
    router.HandleFunc("/pesquisa", controllers.PesquisaFormHandler).Methods("GET")
    
    // Rotas para recuperação de senha.
    router.HandleFunc("/solicitar-recuperacao-senha", controllers.SolicitarRecuperacaoSenha(db)).Methods("POST")
    router.HandleFunc("/solicitar_recuperacao_senha", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/solicitar_recuperacao_senha.html")
    }).Methods("GET")

    router.HandleFunc("/redefinir_senha", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/redefinir_senha.html")
    }).Methods("GET")


    return router
}