package models

import (
    "database/sql"
    "fmt"
)

// Prestador representa um prestador de serviços no condomínio.
type Prestador struct {
    ID              int     `json:"id"`               // ID é o identificador único do prestador.
    NomeEmpresa     string  `json:"nome_empresa"`     // NomeEmpresa é o nome da empresa do prestador.
    TipoServico     string  `json:"tipo_servico"`     // TipoServico é o tipo de serviço prestado.
    Data            string  `json:"data"`             // Data é a data do serviço.
    HoraEntrada     string  `json:"hora_entrada"`     // HoraEntrada é a hora de entrada do prestador.
    HoraSaida       *string `json:"hora_saida"`       // HoraSaida é a hora de saída do prestador.
    NomePrestador   string  `json:"nome_prestador"`   // NomePrestador é o nome do prestador.
    RGCPF           string  `json:"rg_cpf"`           // RGCPF é o RG ou CPF do prestador.
    Telefone        string  `json:"telefone"`         // Telefone é o número de telefone do prestador.
    ContratanteNome string  `json:"contratante_nome"` // ContratanteNome é o nome do contratante do serviço.
    ContratanteTipo string  `json:"contratante_tipo"` // ContratanteTipo é o tipo do contratante (Apartamento ou Condomínio).
    ContratanteApto string  `json:"contratante_apto"` // ContratanteApto é o apartamento do contratante.
    ContratanteBloco string `json:"contratante_bloco"`// ContratanteBloco é o bloco do contratante.
    Autorizou       string  `json:"autorizou"`        // Autorizou é o nome da pessoa que autorizou o serviço.
}

// Search busca prestadores no banco de dados com base em uma query.
// A query é usada para buscar em vários campos do prestador.
func (p *Prestador) Search(db *sql.DB, query string) ([]Prestador, error) {
    query = "%" + query + "%"

    // Query para buscar prestadores no banco de dados.
    rows, err := db.Query(`SELECT id, nome_empresa, tipo_servico, data, hora_entrada, hora_saida, nome_prestador, rg_cpf, telefone, contratante_nome, contratante_tipo, contratante_apto, contratante_bloco, autorizou FROM prestador WHERE nome_empresa LIKE $1 OR tipo_servico LIKE $2 OR nome_prestador LIKE $3 OR rg_cpf LIKE $4 OR telefone LIKE $5 OR contratante_nome LIKE $6 OR contratante_tipo LIKE $7 OR contratante_apto LIKE $8 OR contratante_bloco LIKE $9 OR autorizou LIKE $10`, query, query, query, query, query, query, query, query, query, query)
    if err != nil {
        return nil, fmt.Errorf("erro ao pesquisar prestador: %v", err)
    }
    defer rows.Close()

    var prestadores []Prestador
    for rows.Next() {
        var prestador Prestador
        // Escaneia os resultados da query e preenche a estrutura Prestador.
        if err := rows.Scan(&prestador.ID, &prestador.NomeEmpresa, &prestador.TipoServico, &prestador.Data, &prestador.HoraEntrada, &prestador.HoraSaida, &prestador.NomePrestador, &prestador.RGCPF, &prestador.Telefone, &prestador.ContratanteNome, &prestador.ContratanteTipo, &prestador.ContratanteApto, &prestador.ContratanteBloco, &prestador.Autorizou); err != nil {
            return nil, fmt.Errorf("erro ao escanear prestador: %v", err)
        }
        prestadores = append(prestadores, prestador)
    }
    return prestadores, nil
}

// Create insere um novo prestador no banco de dados.
// Retorna um erro se a inserção falhar.
func (p *Prestador) Create(db *sql.DB) error {
    // Query para inserir um novo prestador e retornar o ID gerado.
    query := `INSERT INTO prestador (nome_empresa, tipo_servico, data, hora_entrada, hora_saida, nome_prestador, rg_cpf, telefone, contratante_nome, contratante_tipo, contratante_apto, contratante_bloco, autorizou) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`
    err := db.QueryRow(query, p.NomeEmpresa, p.TipoServico, p.Data, p.HoraEntrada, p.HoraSaida, p.NomePrestador, p.RGCPF, p.Telefone, p.ContratanteNome, p.ContratanteTipo, p.ContratanteApto, p.ContratanteBloco, p.Autorizou).Scan(&p.ID)
    if err != nil {
        return fmt.Errorf("erro ao criar prestador: %v", err)
    }
    return nil
}