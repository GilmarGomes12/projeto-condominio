package controllers

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"

    "condominio/models"
)

// RegisterFormHandler lida com a exibição do formulário de registro.
func RegisterFormHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Recebida requisição para /register")
    tmpl, err := template.ParseFiles("templates/register.html")
    if err != nil {
        log.Printf("Erro ao carregar template: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    log.Println("Template carregado com sucesso")
    err = tmpl.Execute(w, nil)
    if err != nil {
        log.Printf("Erro ao executar template: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// RegisterUserHandler lida com a criação de um novo usuário.
func RegisterUserHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
            return
        }

        log.Println("Recebida requisição POST para /register")

        // Parseia o formulário
        err := r.ParseForm()
        if err != nil {
            log.Printf("Erro ao parsear formulário: %v", err)
            http.Error(w, "Erro ao processar formulário", http.StatusBadRequest)
            return
        }

        username := r.FormValue("username")
        password := r.FormValue("password")
        role := r.FormValue("role")

        log.Printf("Dados recebidos - Username: %s, Password: %s, Role: %s", username, password, role)

        if username == "" || password == "" {
            log.Println("Nome de usuário e senha são obrigatórios")
            http.Error(w, "Nome de usuário e senha são obrigatórios", http.StatusBadRequest)
            return
        }

        newUser := models.Usuario{
            Username: username,
            Password: password,
        }

        err = newUser.Create(db)
        if err != nil {
            log.Printf("Erro ao criar usuário: %v", err)
            http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
            return
        }

        log.Printf("Usuário criado com ID: %d", newUser.ID)

        // Atribui permissões ao usuário.
        err = models.AssignPermission(db, newUser.ID, role)
        if err != nil {
            log.Printf("Erro ao associar permissão ao usuário: %v", err)
            http.Error(w, "Erro ao associar permissão ao usuário", http.StatusInternalServerError)
            return
        }

        log.Println("Permissão associada com sucesso")

        // Redireciona de volta para a página de registro
        http.Redirect(w, r, "/register", http.StatusSeeOther)
    }
}