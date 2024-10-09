package models

import (
    "database/sql"
    "fmt"
)

// Morador representa um morador do condomínio.
type Morador struct {
    ID          int           `json:"id"`
    Nome        string        `json:"nome"`
    Apartamento string        `json:"apartamento"`
    Bloco       string        `json:"bloco"`
    Telefone1   string        `json:"telefone1"`
    Telefone2   string        `json:"telefone2"`
    Email       string        `json:"email"`
    Email2      string        `json:"email2"`
    Observacao  string        `json:"observacao"`
    Moradores   []MoradorInfo `json:"moradores"`
    Veiculos    []Veiculo     `json:"veiculos"`
}

// MoradorInfo representa informações adicionais sobre um morador.
type MoradorInfo struct {
    Nome           string `json:"nome"`
    DataNascimento string `json:"data_nascimento"`
}

// Veiculo representa um veículo associado a um morador.
type Veiculo struct {
    ID     int    `json:"id"`
    Placa  string `json:"placa"`
    Cor    string `json:"cor"`
    Marca  string `json:"marca"`
    Modelo string `json:"modelo"`
}

// Create insere um novo morador no banco de dados.
// Retorna um erro se a inserção falhar.
func (m *Morador) Create(db *sql.DB) error {
    // Inicia uma transação.
    tx, err := db.Begin()
    if err != nil {
        return fmt.Errorf("erro ao iniciar transação: %v", err)
    }

    // Insere o morador no banco de dados e retorna o ID gerado.
    query := `INSERT INTO moradores (apartamento, bloco, telefone1, telefone2, email, email2, observacao) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
    err = tx.QueryRow(query, m.Apartamento, m.Bloco, m.Telefone1, m.Telefone2, m.Email, m.Email2, m.Observacao).Scan(&m.ID)
    if err != nil {
        tx.Rollback()
        return fmt.Errorf("erro ao criar morador: %v", err)
    }

    // Insere informações adicionais sobre os moradores.
    for _, moradorInfo := range m.Moradores {
        query = `INSERT INTO morador_info (morador_id, nome_completo, data_nascimento) VALUES ($1, $2, $3)`
        _, err = tx.Exec(query, m.ID, moradorInfo.Nome, moradorInfo.DataNascimento)
        if err != nil {
            tx.Rollback()
            return fmt.Errorf("erro ao criar morador_info: %v", err)
        }
    }

    // Insere veículos associados ao morador.
    for _, veiculo := range m.Veiculos {
        query = `INSERT INTO veiculo (placa, cor, marca, modelo) VALUES ($1, $2, $3, $4) RETURNING id`
        err = tx.QueryRow(query, veiculo.Placa, veiculo.Cor, veiculo.Marca, veiculo.Modelo).Scan(&veiculo.ID)
        if err != nil {
            tx.Rollback()
            return fmt.Errorf("erro ao criar veiculo: %v", err)
        }

        query = `INSERT INTO morador_veiculo (morador_id, veiculo_id) VALUES ($1, $2)`
        _, err = tx.Exec(query, m.ID, veiculo.ID)
        if err != nil {
            tx.Rollback()
            return fmt.Errorf("erro ao associar veiculo ao morador: %v", err)
        }
    }

    // Comita a transação.
    err = tx.Commit()
    if err != nil {
        return fmt.Errorf("erro ao commitar transação: %v", err)
    }

    return nil
}

// Search busca moradores no banco de dados com base em uma query.
// A query é usada para buscar em vários campos do morador.
func (m *Morador) Search(db *sql.DB, query string) ([]Morador, error) {
    query = "%" + query + "%"

    // Executa a consulta no banco de dados.
    rows, err := db.Query(`SELECT id, nome, apartamento, bloco FROM moradores WHERE nome LIKE $1 OR apartamento LIKE $2 OR bloco LIKE $3`, query, query, query)
    if err != nil {
        return nil, fmt.Errorf("erro ao pesquisar moradores: %v", err)
    }
    defer rows.Close()

    var moradores []Morador
    for rows.Next() {
        var morador Morador
        // Escaneia os resultados da consulta e preenche a estrutura Morador.
        if err := rows.Scan(&morador.ID, &morador.Nome, &morador.Apartamento, &morador.Bloco); err != nil {
            return nil, fmt.Errorf("erro ao escanear morador: %v", err)
        }
        moradores = append(moradores, morador)
    }
    return moradores, nil
}