package controllers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
    "time"
    "log"
    "condominio/models"
    "condominio/utils"
    "github.com/dgrijalva/jwt-go"
)

// AdminHandler lida com requisições para a administração.
func AdminHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Admin endpoint reached")
    w.Write([]byte("Área de Administração"))
}

// SindicoHandler lida com requisições para o síndico.
func SindicoHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Sindico endpoint reached")
    w.Write([]byte("Área do Síndico"))
}

// PorteiroHandler lida com requisições para o porteiro.
func PorteiroHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Porteiro endpoint reached")
    w.Write([]byte("Área do Porteiro"))
}

// CreateUserHandler lida com a criação de um novo usuário.
func CreateUserHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var usuario struct {
            Username   string   `json:"username"`
            Password   string   `json:"password"`
            Permissoes []string `json:"permissoes"`
        }
        err := json.NewDecoder(r.Body).Decode(&usuario)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        if usuario.Username == "" || usuario.Password == "" {
            http.Error(w, "Nome de usuário e senha são obrigatórios", http.StatusBadRequest)
            return
        }

        newUser := models.Usuario{
            Username: usuario.Username,
            Password: usuario.Password,
        }

        err = newUser.Create(db)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Atribui permissões ao usuário.
        for _, perm := range usuario.Permissoes {
            log.Printf("Atribuindo permissão %s ao usuário ID %d", perm, newUser.ID)
            err = models.AssignPermission(db, newUser.ID, perm)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newUser)
    }
}

// AuthenticateUserHandler lida com a autenticação de um usuário.
func AuthenticateUserHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var creds struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }
        err := json.NewDecoder(r.Body).Decode(&creds)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        usuario, err := models.Authenticate(db, creds.Username, creds.Password)
        if err != nil {
            http.Error(w, err.Error(), http.StatusUnauthorized)
            return
        }

        if usuario == nil {
            http.Error(w, "Usuário ou senha inválidos", http.StatusUnauthorized)
            return
        }

        json.NewEncoder(w).Encode(usuario)
    }
}

// GetUserPermissionsHandler lida com a obtenção das permissões de um usuário.
func GetUserPermissionsHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userIDStr := r.URL.Query().Get("user_id")
        userID, err := strconv.Atoi(userIDStr)
        if err != nil {
            http.Error(w, "ID de usuário inválido", http.StatusBadRequest)
            return
        }

        permissoes, err := models.GetPermissions(db, userID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(permissoes)
    }
}

// SolicitarRecuperacaoSenha lida com a solicitação de recuperação de senha.
func SolicitarRecuperacaoSenha(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var input struct {
            Email string `json:"email" binding:"required,email"`
        }

        err := json.NewDecoder(r.Body).Decode(&input)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        var usuario models.Usuario
        err = db.QueryRow("SELECT id, email FROM usuarios WHERE email = ?", input.Email).Scan(&usuario.ID, &usuario.Email)
        if err != nil {
            http.Error(w, "Usuário não encontrado", http.StatusBadRequest)
            return
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "user_id": usuario.ID,
            "exp":     time.Now().Add(time.Hour * 24).Unix(),
        })

        tokenString, err := token.SignedString([]byte("sua_chave_secreta"))
        if err != nil {
            http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
            return
        }

        // Enviar e-mail com o token
        utils.EnviarEmailRecuperacaoSenha(usuario.Email, tokenString)

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "E-mail de recuperação de senha enviado"})
    }
}

// RedefinirSenha lida com a redefinição de senha.
func RedefinirSenha(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var input struct {
            Token    string `json:"token" binding:"required"`
            Password string `json:"password" binding:"required,min=6"`
        }

        err := json.NewDecoder(r.Body).Decode(&input)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        token, err := jwt.Parse(input.Token, func(token *jwt.Token) (interface{}, error) {
            return []byte("sua_chave_secreta"), nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Token inválido", http.StatusBadRequest)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            http.Error(w, "Token inválido", http.StatusBadRequest)
            return
        }

        userID := int(claims["user_id"].(float64))

        hashedPassword := utils.HashPassword(input.Password)

        _, err = db.Exec("UPDATE usuarios SET password = ? WHERE id = ?", hashedPassword, userID)
        if err != nil {
            http.Error(w, "Erro ao atualizar a senha", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Senha redefinida com sucesso"})
    }
}