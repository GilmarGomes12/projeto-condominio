package models

import (
    "database/sql"
    "fmt"
)

// Visitante representa um visitante no condomínio.
type Visitante struct {
    ID          int    `json:"id"`            // ID é o identificador único do visitante.
    Data        string `json:"data"`          // Data é a data da visita.
    Nome        string `json:"nome_visitante"`// Nome é o nome do visitante.
    Documento   string `json:"rg_cpf"`        // Documento é o RG ou CPF do visitante.
    Visitando   string `json:"visitando"`     // Visitando é a pessoa que o visitante está visitando.
    Apartamento string `json:"apartamento"`   // Apartamento é o número do apartamento visitado.
    Bloco       string `json:"bloco"`         // Bloco é o bloco do apartamento visitado.
    HoraEntrada string `json:"hora_entrada"`  // HoraEntrada é a hora de entrada do visitante.
    HoraSaida   *string `json:"hora_saida"`   // HoraSaida é a hora de saída do visitante.
    Autorizou   string `json:"autorizou"`     // Autorizou é a pessoa que autorizou a visita.
    Vaga        string `json:"vaga"`          // Vaga é a vaga de estacionamento utilizada pelo visitante.
    Placa       string `json:"placa"`         // Placa é a placa do veículo do visitante.
    Marca       string `json:"marca"`         // Marca é a marca do veículo do visitante.
    Modelo      string `json:"modelo"`        // Modelo é o modelo do veículo do visitante.
    Cor         string `json:"cor"`           // Cor é a cor do veículo do visitante.
}

// Search busca visitantes no banco de dados com base em uma query.
// A query é usada para buscar em vários campos do visitante.
func (v *Visitante) Search(db *sql.DB, query string) ([]Visitante, error) {
    query = "%" + query + "%"

    // Executa a query para buscar visitantes no banco de dados.
    rows, err := db.Query(`SELECT id, data, nome_visitante, rg_cpf, visitando, apartamento, bloco, hora_entrada, hora_saida, autorizou, vaga, placa, marca, modelo, cor FROM visitante WHERE data LIKE $1 OR nome_visitante LIKE $2 OR rg_cpf LIKE $3 OR visitando LIKE $4 OR apartamento LIKE $5 OR bloco LIKE $6 OR hora_entrada LIKE $7 OR hora_saida LIKE $8 OR autorizou LIKE $9 OR vaga LIKE $10 OR placa LIKE $11 OR marca LIKE $12 OR modelo LIKE $13 OR cor LIKE $14`, query, query, query, query, query, query, query, query, query, query, query, query, query, query)
    if err != nil {
        return nil, fmt.Errorf("erro ao pesquisar visitante: %v", err)
    }
    defer rows.Close()

    var visitantes []Visitante
    for rows.Next() {
        var visitante Visitante
        // Escaneia os resultados da query e preenche a estrutura Visitante.
        if err := rows.Scan(&visitante.ID, &visitante.Data, &visitante.Nome, &visitante.Documento, &visitante.Visitando, &visitante.Apartamento, &visitante.Bloco, &visitante.HoraEntrada, &visitante.HoraSaida, &visitante.Autorizou, &visitante.Vaga, &visitante.Placa, &visitante.Marca, &visitante.Modelo, &visitante.Cor); err != nil {
            return nil, fmt.Errorf("erro ao escanear visitante: %v", err)
        }
        visitantes = append(visitantes, visitante)
    }
    return visitantes, nil
}

// Create insere um novo visitante no banco de dados.
// Retorna um erro se a inserção falhar.
func (v *Visitante) Create(db *sql.DB) error {
    // Query para inserir um novo visitante e retornar o ID gerado.
    query := `INSERT INTO visitante (data, nome_visitante, rg_cpf, visitando, apartamento, bloco, hora_entrada, hora_saida, autorizou, vaga, placa, marca, modelo, cor) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id`
    err := db.QueryRow(query, v.Data, v.Nome, v.Documento, v.Visitando, v.Apartamento, v.Bloco, v.HoraEntrada, v.HoraSaida, v.Autorizou, v.Vaga, v.Placa, v.Marca, v.Modelo, v.Cor).Scan(&v.ID)
    if err != nil {
        return fmt.Errorf("erro ao criar visitante: %v", err)
    }
    return nil
}