package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Funcionarios representa um funcionário do condomínio.
type Funcionarios struct {
    ID             int       `json:"id"`              // ID é o identificador único do funcionário.
    Nome           string    `json:"nome"`            // Nome é o nome do funcionário.
    Endereco       string    `json:"endereco"`        // Endereco é o endereço do funcionário.
    Bairro         string    `json:"bairro"`          // Bairro é o bairro do funcionário.
    CEP            string    `json:"cep"`             // CEP é o código postal do funcionário.
    Cidade         string    `json:"cidade"`          // Cidade é a cidade do funcionário.
    UF             string    `json:"uf"`              // UF é a unidade federativa do funcionário.
    Telefone       string    `json:"telefone"`        // Telefone é o telefone do funcionário.
    Celular        string    `json:"celular"`         // Celular é o celular do funcionário.
    Email          string    `json:"email"`           // Email é o email do funcionário.
    FuncaoCargo    string    `json:"funcao_cargo"`    // FuncaoCargo é a função ou cargo do funcionário.
    HorarioTrabalho string   `json:"horario_trabalho"` // HorarioTrabalho é o horário de trabalho do funcionário.
    AdmitidoEm     time.Time `json:"admitido_em"`      // AdmitidoEm é a data de admissão do funcionário.
    Observacoes    string    `json:"observacoes"`     // Observacoes são as observações sobre o funcionário.
}

// Create insere um novo funcionário no banco de dados.
func (f *Funcionarios) Create(db *sql.DB) error {
    query := `INSERT INTO funcionario (nome, endereco, bairro, cep, cidade, uf, telefone, celular, email, funcao_cargo, horario_trabalho, admitido_em, observacoes)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`

              log.Printf("Valores para inserção: %+v", f)          
    err := db.QueryRow(query, f.Nome, f.Endereco, f.Bairro, f.CEP, f.Cidade, f.UF, f.Telefone, f.Celular, f.Email, f.FuncaoCargo, f.HorarioTrabalho, f.AdmitidoEm, f.Observacoes).Scan(&f.ID)
    if err != nil {
        return fmt.Errorf("erro ao criar funcionario: %v", err)
    }
    return nil
}

// GetAll retorna todos os funcionários do banco de dados.
func (f *Funcionarios) GetAll(db *sql.DB) ([]Funcionarios, error) {
rows, err := db.Query(`SELECT id, nome, endereco, bairro, cep, cidade, uf, telefone, celular, email, funcao_cargo, horario_trabalho, admitido_em, observacoes FROM funcionario`)
if err != nil {
    return nil, fmt.Errorf("erro ao buscar funcionarios: %v", err)
}
defer rows.Close()

    var funcionarios []Funcionarios
    for rows.Next() {
        var funcionario Funcionarios
        if err := rows.Scan(&funcionario.ID, &funcionario.Nome, &funcionario.Endereco, &funcionario.Bairro, &funcionario.CEP, &funcionario.Cidade, &funcionario.UF, &funcionario.Telefone, &funcionario.Celular, &funcionario.Email, &funcionario.FuncaoCargo, &funcionario.HorarioTrabalho, &funcionario.AdmitidoEm, &funcionario.Observacoes); err != nil {
            return nil, fmt.Errorf("erro ao escanear funcionario: %v", err)
        }
        funcionarios = append(funcionarios, funcionario)
    }
    return funcionarios, nil

}

// Search method for Funcionarios to satisfy the interface used in the controller
func (f *Funcionarios) Search(db *sql.DB, query string) (interface{}, error) {

    // Implement the search logic here

    return nil, nil
}