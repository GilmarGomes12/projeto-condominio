package models

import (
    "database/sql"
    "fmt"
    "time"
)

// Ocorrencia representa uma ocorrência registrada no condomínio.
type Ocorrencia struct {
    NumeroOcorrencia    string    `json:"numero_ocorrencia"`
    DataOcorrencia      time.Time `json:"data_ocorrencia" db:"data_ocorrencia"`
    NomeFuncionario     string    `json:"nome_funcionario"`
    FuncaoFuncionario   string    `json:"funcao_funcionario"`
    HoraRegistro        string    `json:"hora_registro"`
    DataRegistro        time.Time `json:"data_registro" db:"data_registro"`
    UnidadeEnvolvida    string    `json:"unidade_envolvida"`
    Bloco               string    `json:"bloco"` 
    AutorOcorrencia     string    `json:"autor_ocorrencia"`
    DescricaoOcorrencia string    `json:"descricao_ocorrencia"`
}

// Create insere uma nova ocorrência no banco de dados.
func (o *Ocorrencia) Create(db *sql.DB) error {
    query := `INSERT INTO ocorrencia (numero_ocorrencia, data_ocorrencia, nome_funcionario, funcao_funcionario, hora_registro, data_registro, unidade_envolvida, bloco, autor_ocorrencia, descricao_ocorrencia) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
    _, err := db.Exec(query, o.NumeroOcorrencia, o.DataOcorrencia, o.NomeFuncionario, o.FuncaoFuncionario, o.HoraRegistro, o.DataRegistro, o.UnidadeEnvolvida, o.Bloco, o.AutorOcorrencia, o.DescricaoOcorrencia)
    if err != nil {
        return fmt.Errorf("erro ao criar ocorrência: %v", err)
    }
    return nil
}

// Search busca ocorrências no banco de dados com base em uma query.
func (o *Ocorrencia) Search(db *sql.DB, query string) ([]Ocorrencia, error) {
    query = "%" + query + "%"

    rows, err := db.Query(`SELECT numero_ocorrencia, data_ocorrencia, nome_funcionario, funcao_funcionario, hora_registro, data_registro, unidade_envolvida, autor_ocorrencia, descricao_ocorrencia FROM ocorrencia WHERE nome_funcionario LIKE $1 OR unidade_envolvida LIKE $2 OR autor_ocorrencia LIKE $3`, query, query, query)
    if err != nil {
        return nil, fmt.Errorf("erro ao pesquisar ocorrências: %v", err)
    }
    defer rows.Close()

    var ocorrencias []Ocorrencia
    for rows.Next() {
        var ocorrencia Ocorrencia
        if err := rows.Scan(&ocorrencia.NumeroOcorrencia, &ocorrencia.DataOcorrencia, &ocorrencia.NomeFuncionario, &ocorrencia.FuncaoFuncionario, &ocorrencia.HoraRegistro, &ocorrencia.DataRegistro, &ocorrencia.UnidadeEnvolvida, &ocorrencia.AutorOcorrencia, &ocorrencia.DescricaoOcorrencia); err != nil {
            return nil, fmt.Errorf("erro ao escanear ocorrência: %v", err)
        }
        ocorrencias = append(ocorrencias, ocorrencia)
    }
    return ocorrencias, nil
}