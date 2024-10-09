package models

import (
    "database/sql"
    "fmt"
)

// Agendamento representa um agendamento no condomínio.
type Agendamento struct {
    ID          int    `json:"id"`           // ID é o identificador único do agendamento.
    Nome        string `json:"nome_morador"` // Nome é o nome do morador que fez o agendamento.
    Apartamento string `json:"apartamento"`  // Apartamento é o número do apartamento do morador.
    Bloco       string `json:"bloco"`        // Bloco é o bloco do apartamento do morador.
    Local       string `json:"local"`        // Local é o local onde o agendamento será realizado.
    Dia         int    `json:"dia"`          // Dia é o dia do agendamento.
    Mes         int    `json:"mes"`          // Mes é o mês do agendamento.
    Ano         int    `json:"ano"`          // Ano é o ano do agendamento.
    Periodo     string `json:"periodo"`      // Periodo é o período do agendamento (manhã, tarde, noite).
    Funcionario string `json:"funcionario"`  // Funcionario é o funcionário responsável pelo agendamento.
    Observacoes string `json:"observacoes"`  // Observacoes são observações adicionais sobre o agendamento.
    Convidados  string `json:"convidados"`   // Convidados são os convidados do morador para o agendamento.
}

// Search busca agendamentos no banco de dados com base em uma query.
// A query é usada para buscar em vários campos do agendamento.
func (a *Agendamento) Search(db *sql.DB, query string) ([]Agendamento, error) {
    query = "%" + query + "%"

    // Executa a consulta no banco de dados.
    rows, err := db.Query(`SELECT id, nome_morador, apartamento, bloco, local, dia, mes, ano, periodo, funcionario, observacoes, convidados FROM agendamento WHERE nome_morador LIKE $1 OR apartamento LIKE $2 OR bloco LIKE $3 OR local LIKE $4 OR data LIKE $5 OR periodo LIKE $6`, query, query, query, query, query, query)
    if err != nil {
        return nil, fmt.Errorf("erro ao pesquisar agendamentos: %v", err)
    }
    defer rows.Close()

    var agendamentos []Agendamento
    for rows.Next() {
        var agendamento Agendamento
        // Escaneia os resultados da consulta e preenche a estrutura Agendamento.
        if err := rows.Scan(&agendamento.ID, &agendamento.Nome, &agendamento.Apartamento, &agendamento.Bloco, &agendamento.Local, &agendamento.Dia, &agendamento.Mes, &agendamento.Ano, &agendamento.Periodo, &agendamento.Funcionario, &agendamento.Observacoes, &agendamento.Convidados); err != nil {
            return nil, fmt.Errorf("erro ao escanear agendamento: %v", err)
        }
        agendamentos = append(agendamentos, agendamento)
    }
    return agendamentos, nil
}

// Create insere um novo agendamento no banco de dados.
func (a *Agendamento) Create(db *sql.DB) error {
    // Executa a inserção no banco de dados e retorna o ID gerado.
    query := `INSERT INTO agendamento (nome_morador, apartamento, bloco, local, dia, mes, ano, periodo, funcionario, observacoes, convidados) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
    err := db.QueryRow(query, a.Nome, a.Apartamento, a.Bloco, a.Local, a.Dia, a.Mes, a.Ano, a.Periodo, a.Funcionario, a.Observacoes, a.Convidados).Scan(&a.ID)
    if err != nil {
        return fmt.Errorf("erro ao criar agendamento: %v", err)
    }
    return nil
}