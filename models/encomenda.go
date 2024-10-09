package models

import (
    "database/sql"
    "fmt"
    "log"
    "time"
    "github.com/google/uuid"
)

// Encomenda representa uma encomenda recebida no condomínio.
type Encomenda struct {
    NumeroProtocolo      string    `json:"numero_protocolo"`      // NumeroProtocolo é o identificador único da encomenda.
    DataHoraRecebimento  time.Time `json:"data_hora_recebimento"` // DataHoraRecebimento é a data e hora em que a encomenda foi recebida.
    NomeDestinatario     string    `json:"nome_destinatario"`     // NomeDestinatario é o nome do destinatário da encomenda.
    Apartamento          string    `json:"apartamento"`           // Apartamento é o número do apartamento do destinatário.
    Bloco                string    `json:"bloco"`                 // Bloco é o bloco do apartamento do destinatário.
    NumeroRastreamento   string    `json:"numero_rastreamento"`   // NumeroRastreamento é o número de rastreamento da encomenda.
    TipoEncomenda        string    `json:"tipo_encomenda"`        // TipoEncomenda é o tipo da encomenda.
    DescricaoEncomenda   string    `json:"descricao_encomenda"`   // DescricaoEncomenda é a descrição da encomenda.
    EmpresaEntrega       string    `json:"empresa_entrega"`       // EmpresaEntrega é a empresa responsável pela entrega.
    Observacoes          string    `json:"observacoes"`           // Observacoes são observações adicionais sobre a encomenda.
    NomeEntregador       string    `json:"nome_entregador"`       // NomeEntregador é o nome do entregador.
    CpfRgEntregador      string    `json:"cpf_rg_entregador"`     // CpfRgEntregador é o CPF ou RG do entregador.
    NomePorteiro         string    `json:"nome_porteiro"`         // NomePorteiro é o nome do porteiro que recebeu a encomenda.
    NomeRetirou          string    `json:"nome_retirou"`          // NomeRetirou é o nome da pessoa que retirou a encomenda.
}

// Create insere uma nova encomenda no banco de dados.
// Gera um número de protocolo único e define a data e hora de recebimento.
func (e *Encomenda) Create(db *sql.DB) error {
    e.NumeroProtocolo = uuid.New().String()
    e.DataHoraRecebimento = time.Now()

    query := `INSERT INTO encomenda (numero_protocolo, data_hora_recebimento, nome_destinatario, apartamento, bloco, numero_rastreamento, tipo_encomenda, descricao_encomenda, empresa_entrega, observacoes, nome_entregador, cpf_rg_entregador, nome_porteiro, nome_retirou) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
    _, err := db.Exec(query, e.NumeroProtocolo, e.DataHoraRecebimento, e.NomeDestinatario, e.Apartamento, e.Bloco, e.NumeroRastreamento, e.TipoEncomenda, e.DescricaoEncomenda, e.EmpresaEntrega, e.Observacoes, e.NomeEntregador, e.CpfRgEntregador, e.NomePorteiro, e.NomeRetirou)
    if err != nil {
        return fmt.Errorf("erro ao criar encomenda: %v", err)
    }
    return nil
}

// Search busca encomendas no banco de dados com base em uma query.
// A query é usada para buscar em vários campos da encomenda.
func (e *Encomenda) Search(db *sql.DB, query string) ([]Encomenda, error) {
    query = "%" + query + "%"

    // Executa a consulta no banco de dados.
    rows, err := db.Query(`SELECT numero_protocolo, data_hora_recebimento, nome_destinatario, apartamento, bloco, numero_rastreamento, tipo_encomenda, descricao_encomenda, empresa_entrega, observacoes, nome_entregador, cpf_rg_entregador, nome_porteiro, nome_retirou FROM encomenda WHERE nome_destinatario LIKE $1 OR apartamento LIKE $2 OR bloco LIKE $3`, query, query, query)
    if err != nil {
        return nil, fmt.Errorf("erro ao pesquisar encomendas: %v", err)
    }
    defer rows.Close()

    var encomendas []Encomenda
    for rows.Next() {
        var encomenda Encomenda
        // Escaneia os resultados da consulta e preenche a estrutura Encomenda.
        if err := rows.Scan(&encomenda.NumeroProtocolo, &encomenda.DataHoraRecebimento, &encomenda.NomeDestinatario, &encomenda.Apartamento, &encomenda.Bloco, &encomenda.NumeroRastreamento, &encomenda.TipoEncomenda, &encomenda.DescricaoEncomenda, &encomenda.EmpresaEntrega, &encomenda.Observacoes, &encomenda.NomeEntregador, &encomenda.CpfRgEntregador, &encomenda.NomePorteiro, &encomenda.NomeRetirou); err != nil {
            return nil, fmt.Errorf("erro ao escanear encomenda: %v", err)
        }
        encomendas = append(encomendas, encomenda)
    }
    return encomendas, nil
}

// GetAll retorna todas as encomendas do banco de dados.
func (e *Encomenda) GetAll(db *sql.DB) ([]Encomenda, error) {
    // Executa a consulta no banco de dados.
    rows, err := db.Query(`SELECT numero_protocolo, data_hora_recebimento, nome_destinatario, apartamento, bloco, numero_rastreamento, tipo_encomenda, descricao_encomenda, empresa_entrega, observacoes, nome_entregador, cpf_rg_entregador, nome_porteiro, nome_retirou FROM encomenda`)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar todas as encomendas: %v", err)
    }
    defer rows.Close()

    var encomendas []Encomenda
    for rows.Next() {
        var encomenda Encomenda
        // Escaneia os resultados da consulta e preenche a estrutura Encomenda.
        if err := rows.Scan(&encomenda.NumeroProtocolo, &encomenda.DataHoraRecebimento, &encomenda.NomeDestinatario, &encomenda.Apartamento, &encomenda.Bloco, &encomenda.NumeroRastreamento, &encomenda.TipoEncomenda, &encomenda.DescricaoEncomenda, &encomenda.EmpresaEntrega, &encomenda.Observacoes, &encomenda.NomeEntregador, &encomenda.CpfRgEntregador, &encomenda.NomePorteiro, &encomenda.NomeRetirou); err != nil {
            log.Printf("erro ao escanear encomenda: %v", err)
            return nil, fmt.Errorf("erro ao escanear encomenda: %v", err)
        }
        encomendas = append(encomendas, encomenda)
    }
    log.Printf("encomendas encontradas: %v", encomendas)
    return encomendas, nil
}