package models

import (
    "database/sql"
    "fmt"
    "time"
)

// AgendamentoMudanca representa um agendamento de mudança no condomínio.
type Mudanca struct {
    ID                int       `json:"id"`                // ID é o identificador único do agendamento.
    DataMudanca       time.Time `json:"data_mudanca"`      // DataMudanca é a data da mudança.
    ResponsavelNome   string    `json:"responsavel_nome"`  // ResponsavelNome é o nome do responsável pela mudança.
    ResponsavelApto   string    `json:"responsavel_apto"`  // ResponsavelApto é o apartamento do responsável pela mudança.
    ResponsavelBloco  string    `json:"responsavel_bloco"` // ResponsavelBloco é o bloco do responsável pela mudança.
    Horario           string    `json:"horario"`           // Horario é o horário da mudança.
    NomeEmpresa       string    `json:"nome_empresa"`      // NomeEmpresa é o nome da empresa responsável pela mudança.
    UsoElevador       bool      `json:"uso_elevador"`      // UsoElevador indica se o elevador será usado na mudança.
    UsoEscada         bool      `json:"uso_escada"`        // UsoEscada indica se a escada será usada na mudança.
    Iscar             bool      `json:"iscar"`             // Iscar indica se um carro será usado na mudança.
}

// Create insere um novo agendamento de mudança no banco de dados.
func (a *Mudanca) Create(db *sql.DB) error {
    query := `INSERT INTO mudanca (data_mudanca, responsavel_nome, responsavel_apto, responsavel_bloco, horario, nome_empresa, uso_elevador, uso_escada, iscar) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
    err := db.QueryRow(query, a.DataMudanca, a.ResponsavelNome, a.ResponsavelApto, a.ResponsavelBloco, a.Horario, a.NomeEmpresa, a.UsoElevador, a.UsoEscada, a.Iscar).Scan(&a.ID)
    if err != nil {
        return fmt.Errorf("erro ao criar agendamento de mudança: %v", err)
    }
    return nil
}

// GetAll retorna todos os agendamentos de mudança do banco de dados.
func (a *Mudanca) GetAll(db *sql.DB) ([]Mudanca, error) {
    rows, err := db.Query(`SELECT id, data_mudanca, responsavel_nome, responsavel_apto, responsavel_bloco, horario, nome_empresa, uso_elevador, uso_escada, iscar FROM agendamento_mudanca`)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar agendamentos de mudança: %v", err)
    }

    defer rows.Close()

    var agendamentos []Mudanca

    return agendamentos, nil
}

// Search method for Mudanca type
func (m *Mudanca) Search(db *sql.DB, query string) ([]Mudanca, error) {

    // Example query, adjust as needed
    rows, err := db.Query("SELECT * FROM mudancas WHERE field LIKE ?", "%"+query+"%")
    if err != nil {
        return nil, err
    }
    
    defer rows.Close()

    var agendamentos []Mudanca
    for rows.Next() {
        var agendamento Mudanca
        if err := rows.Scan(&agendamento.ID, &agendamento.DataMudanca, &agendamento.ResponsavelNome, &agendamento.ResponsavelApto, &agendamento.ResponsavelBloco, &agendamento.Horario, &agendamento.NomeEmpresa, &agendamento.UsoElevador, &agendamento.UsoEscada, &agendamento.Iscar); err != nil {
            return nil, fmt.Errorf("erro ao escanear agendamento de mudança: %v", err)
        }
        agendamentos = append(agendamentos, agendamento)
    }
    return agendamentos, nil
}