package controllers

import (
    "condominio/models"
    "database/sql"
    "html/template"
    "log"
    "net/http"
    "github.com/gorilla/sessions"
)

// Variáveis globais para o template de login e o armazenamento de sessões.
var (
    tmpl = template.Must(template.ParseFiles("templates/login.html")) // Carrega o template de login.
    store = sessions.NewCookieStore([]byte("super-secret-key"))       // Inicializa o armazenamento de sessões com uma chave secreta.
)

// LoginHandler lida com as requisições HTTP para a página de login.
func LoginHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            // Renderiza o template de login para requisições GET.
            err := tmpl.Execute(w, nil)
            if err != nil {
                log.Fatalf("Erro ao carregar o template: %v", err)
            }
        } else if r.Method == http.MethodPost {
            // Obtém os valores do formulário de login.
            username := r.FormValue("username")
            password := r.FormValue("password")

            // Autentica o usuário com base nas credenciais fornecidas.
            usuario, err := models.Authenticate(db, username, password)
            if err != nil || usuario == nil {
                // Redireciona para a página de login se a autenticação falhar.
                http.Redirect(w, r, "/login", http.StatusSeeOther)
                return
            }

            // Cria uma nova sessão e armazena o ID do usuário.
            session, _ := store.Get(r, "session-name")
            session.Values["userID"] = usuario.ID
            session.Save(r, w)

            // Redireciona para a página inicial se as credenciais forem válidas.
            http.Redirect(w, r, "/menu", http.StatusSeeOther)
        }
    }
}