package models

import (
    "database/sql"
    "log"
)

// Usuario representa um usuário no sistema.
type Usuario struct {
    ID       int    // ID é o identificador único do usuário.
    Username string // Username é o nome de usuário.
    Password string // Password é a senha do usuário.
    Email    string `json:"email"`
}

// Permissao representa uma permissão associada a um usuário.
type Permissao struct {
    Nome string // Nome é o nome da permissão.
}

// Create insere um novo usuário no banco de dados.
func (u *Usuario) Create(db *sql.DB) error {
    query := "INSERT INTO usuarios (username, password) VALUES ($1, $2) RETURNING id"
    log.Printf("Executando query: %s com valores: %s, %s", query, u.Username, u.Password)
    err := db.QueryRow(query, u.Username, u.Password).Scan(&u.ID)
    if err != nil {
        log.Printf("Erro ao executar query: %v", err)
        return err
    }
    return nil
}

// GetPermissions retorna as permissões de um usuário.
func GetPermissions(db *sql.DB, userID int) ([]Permissao, error) {
    query := "SELECT permissao FROM permissoes WHERE user_id = $1"
    rows, err := db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var permissoes []Permissao
    for rows.Next() {
        var p Permissao
        if err := rows.Scan(&p.Nome); err != nil {
            log.Println("Error scanning permission:", err)
            return nil, err
        }
        permissoes = append(permissoes, p)
    }

    if err := rows.Err(); err != nil {
        log.Println("Error with rows:", err)
        return nil, err
    }

    return permissoes, nil
}

// AssignPermission associa uma permissão a um usuário.
func AssignPermission(db *sql.DB, userID int, permission string) error {
    query := "INSERT INTO permissoes (user_id, permissao) VALUES ($1, $2)"
    log.Printf("Executando query: %s com valores: %d, %s", query, userID, permission)
    _, err := db.Exec(query, userID, permission)
    if err != nil {
        log.Println("Error assigning permission:", err)
        return err
    }
    return nil
}

// Authenticate verifica as credenciais do usuário.
func Authenticate(db *sql.DB, username, password string) (*Usuario, error) {
    query := "SELECT id, username, password FROM usuarios WHERE username = $1 AND password = $2"
    row := db.QueryRow(query, username, password)

    var u Usuario
    err := row.Scan(&u.ID, &u.Username, &u.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Println("Usuário não encontrado ou senha incorreta")
            return nil, nil
        }
        log.Println("Erro ao autenticar usuário:", err)
        return nil, err
    }

    return &u, nil
}