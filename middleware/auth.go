package middleware

import (
    "condominio/models"
    "context"
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/sessions"
)

// Inicializa o armazenamento de sessões com uma chave secreta.
var store = sessions.NewCookieStore([]byte("super-secret-key"))

type contextKey string

const userIDKey contextKey = "userID"

// RequirePermission verifica se o usuário tem a permissão necessária para acessar um recurso.
// Se o usuário não estiver autenticado ou não tiver a permissão necessária, retorna um erro 403.
func RequirePermission(permission string, db *sql.DB, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Obtém a sessão do usuário.
        session, _ := store.Get(r, "session-name")
        userID, ok := session.Values["userID"]
        if !ok {
            log.Println("Usuário não autenticado")
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        // Verifica o tipo de userID e converte para int.
        userIDInt, err := strconv.Atoi(fmt.Sprintf("%v", userID))
        if err != nil {
            log.Println("Erro ao converter userID para int:", err)
            http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
            return
        }

        // Obtém as permissões do usuário a partir do banco de dados.
        log.Println("Obtendo permissões para o usuário ID:", userIDInt)
        permissoes, err := models.GetPermissions(db, userIDInt)
        if err != nil {
            log.Println("Erro ao obter permissões do usuário:", err)
            http.Error(w, "Erro ao obter permissões do usuário", http.StatusInternalServerError)
            return
        }

        // Verifica se o usuário tem a permissão necessária.
        hasPermission := false
        for _, p := range permissoes {
            if p.Nome == permission {
                hasPermission = true
                break
            }
        }

        if !hasPermission {
            log.Println("Usuário não tem permissão:", permission)
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        // Adiciona o ID do usuário ao contexto da requisição e chama o próximo handler.
        log.Println("Permissão concedida para o usuário ID:", userIDInt)
        ctx := context.WithValue(r.Context(), userIDKey, strconv.Itoa(userIDInt))
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}

// WithUserID adiciona o ID do usuário ao contexto da requisição.
// Se o usuário não estiver autenticado, redireciona para a página de login.
func WithUserID(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Obtém a sessão do usuário.
        session, _ := store.Get(r, "session-name")
        userID, ok := session.Values["userID"]
        if !ok {
            log.Println("Usuário não autenticado, redirecionando para login")
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        // Verifica o tipo de userID e converte para int.
        userIDInt, err := strconv.Atoi(fmt.Sprintf("%v", userID))
        if err != nil {
            log.Println("Erro ao converter userID para int:", err)
            http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
            return
        }

        // Adiciona o ID do usuário ao contexto da requisição e chama o próximo handler.
        log.Println("Usuário autenticado com ID:", userIDInt)
        ctx := context.WithValue(r.Context(), userIDKey, strconv.Itoa(userIDInt))
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}