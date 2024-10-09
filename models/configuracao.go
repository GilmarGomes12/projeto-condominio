package models

import (
    "database/sql"
    "fmt"
)

// Configuracao representa uma configuração do sistema.
type Configuracao struct {
    ID    int    `json:"id"`    // ID é o identificador único da configuração.
    Nome  string `json:"nome"`  // Nome é o nome da configuração.
    Senha string `json:"senha"` // Senha é a senha associada à configuração.
}

// Create insere uma nova configuração no banco de dados.
// Retorna um erro se a inserção falhar.
func (c *Configuracao) Create(db *sql.DB) error {
    // Query para inserir uma nova configuração e retornar o ID gerado.
    query := `INSERT INTO configuracoes (nome, senha) VALUES ($1, $2) RETURNING id`
    err := db.QueryRow(query, c.Nome, c.Senha).Scan(&c.ID)
    if err != nil {
        return fmt.Errorf("erro ao criar configuracao: %v", err)
    }
    return nil
}

// GetAll retorna todas as configurações do banco de dados.
// Retorna um slice de Configuracao e um erro, se houver.
func (c *Configuracao) GetAll(db *sql.DB) ([]Configuracao, error) {
    // Query para buscar todas as configurações.
    rows, err := db.Query(`SELECT id, nome, senha FROM configuracoes`)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar configuracoes: %v", err)
    }
    defer rows.Close()

    var configuracoes []Configuracao
    for rows.Next() {
        var configuracao Configuracao
        // Escaneia os resultados da query e preenche a estrutura Configuracao.
        if err := rows.Scan(&configuracao.ID, &configuracao.Nome, &configuracao.Senha); err != nil {
            return nil, fmt.Errorf("erro ao escanear configuracao: %v", err)
        }
        configuracoes = append(configuracoes, configuracao)
    }
    return configuracoes, nil
}