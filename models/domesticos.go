package models

import (
    "database/sql"
    "fmt"
)

// FuncionarioDomestico representa um funcionário doméstico no condomínio.
type Domesticos struct {
    ID          int    `json:"id"`          // ID é o identificador único do funcionário doméstico.
    Nome        string `json:"nome"`        // Nome é o nome do funcionário doméstico.
    Apartamento string `json:"apartamento"` // Apartamento é o número do apartamento onde o funcionário trabalha.
    Bloco       string `json:"bloco"`       // Bloco é o bloco do apartamento onde o funcionário trabalha.
    Funcao      string `json:"funcao"`      // Funcao é a função do funcionário doméstico.
    Horario     string `json:"horario"`     // Horario é o horário de trabalho do funcionário doméstico.
    Telefone    string `json:"telefone"`    // Telefone é o número de telefone do funcionário doméstico.
}

// Create insere um novo funcionário doméstico no banco de dados.
// Retorna um erro se a inserção falhar.
func (f *Domesticos) Create(db *sql.DB) error {
    // Query para inserir um novo funcionário doméstico e retornar o ID gerado.
    query := `INSERT INTO domesticos (nome, apartamento, bloco, funcao, horario, telefone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
    err := db.QueryRow(query, f.Nome, f.Apartamento, f.Bloco, f.Funcao, f.Horario, f.Telefone).Scan(&f.ID)
    if err != nil {
        return fmt.Errorf("erro ao criar funcionario domestico: %v", err)
    }
    return nil
}

// GetAll retorna todos os funcionários domésticos do banco de dados.
// Retorna um slice de FuncionarioDomestico e um erro, se houver.
func (f *Domesticos) GetAll(db *sql.DB) ([]Domesticos, error) {
    // Query para buscar todos os funcionários domésticos.
    rows, err := db.Query(`SELECT id, nome, apartamento, bloco, funcao, horario, telefone FROM domesticos`)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar funcionarios domesticos: %v", err)
    }
    defer rows.Close()

    var funcionarios []Domesticos
    for rows.Next() {
        var funcionario Domesticos
        // Escaneia os resultados da query e preenche a estrutura FuncionarioDomestico.
        if err := rows.Scan(&funcionario.ID, &funcionario.Nome, &funcionario.Apartamento, &funcionario.Bloco, &funcionario.Funcao, &funcionario.Horario, &funcionario.Telefone); err != nil {
            return nil, fmt.Errorf("erro ao escanear funcionario domestico: %v", err)
        }
        funcionarios = append(funcionarios, funcionario)
    }

        return funcionarios, nil
    }
    
// Search busca funcionários domésticos pelo nome no banco de dados.
// Retorna um slice de Domesticos e um erro, se houver.
func (f *Domesticos) Search(db *sql.DB, query string) ([]Domesticos, error) {
    // Query para buscar funcionários domésticos pelo nome.
    rows, err := db.Query("SELECT id, nome, apartamento, bloco, funcao, horario, telefone FROM domesticos WHERE nome LIKE $1", "%"+query+"%")
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar funcionarios domesticos: %v", err)
    }
    defer rows.Close()

    var results []Domesticos
    for rows.Next() {
        var funcionario Domesticos
        if err := rows.Scan(&funcionario.ID, &funcionario.Nome, &funcionario.Apartamento, &funcionario.Bloco, &funcionario.Funcao, &funcionario.Horario, &funcionario.Telefone); err != nil {
            return nil, fmt.Errorf("erro ao escanear funcionario domestico: %v", err)
        }
        results = append(results, funcionario)
    }

    return results, nil
}