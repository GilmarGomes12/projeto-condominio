package controllers

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"
    "condominio/models"
)

// DomesticosHandler lida com as requisições HTTP para gerenciar funcionários domésticos.
// Suporta métodos POST para criar um novo funcionário e GET para renderizar a página de funcionários.
func DomesticosHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            // Trata requisições POST para criar um novo funcionário doméstico.
            if err := r.ParseForm(); err != nil {
                log.Printf("Erro ao parsear o formulário: %v", err)
                http.Error(w, "Erro ao parsear o formulário", http.StatusBadRequest)
                return
            }

            // Extrai os valores dos campos do formulário.
            nome := r.FormValue("nome")
            apartamento := r.FormValue("apartamento")
            bloco := r.FormValue("bloco")
            funcao := r.FormValue("funcao")
            horario := r.FormValue("horario")
            telefone := r.FormValue("telefone")

            // Verifica se todos os campos obrigatórios estão presentes.
            if nome == "" || apartamento == "" || bloco == "" || funcao == "" || horario == "" || telefone == "" {
                http.Error(w, "Campos obrigatórios faltando", http.StatusBadRequest)
                return
            }

            // Cria um novo funcionário doméstico.
            funcionario := models.Domesticos{
                Nome:        nome,
                Apartamento: apartamento,
                Bloco:       bloco,
                Funcao:      funcao,
                Horario:     horario,
                Telefone:    telefone,
            }

            // Tenta criar o novo funcionário no banco de dados.
            if err := funcionario.Create(db); err != nil {
                log.Printf("Erro ao criar funcionário: %v", err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Redireciona para a página de sucesso ou para a mesma página para limpar o formulário.
            http.Redirect(w, r, "/domesticos", http.StatusSeeOther)

        case "GET":
            // Carrega o template para a página de listagem/cadastro dos funcionários.
            tmpl, err := template.ParseFiles("templates/domesticos.html")
            if err != nil {
                log.Printf("Erro ao carregar o template: %v", err)
                http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
                return
            }
            err = tmpl.Execute(w, nil)
            if err != nil {
                log.Printf("Erro ao executar o template: %v", err)
                http.Error(w, "Erro ao executar o template", http.StatusInternalServerError)
            }

        default:
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        }
    }
}
