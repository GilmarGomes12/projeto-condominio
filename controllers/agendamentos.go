package controllers

import (
    "condominio/models"
    "database/sql"
    "html/template"
    "log"
    "net/http"
    "strconv"
)

// AgendamentosFormHandler lida com requisições HTTP para exibir o formulário de agendamentos.
// Suporta apenas o método GET para renderizar a página de agendamentos.
func AgendamentosFormHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/agendamentos.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

// AgendamentosHandler lida com requisições HTTP para criar novos agendamentos.
// Suporta apenas o método POST para criar um novo agendamento.
func AgendamentosHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            // Parseia o formulário
            err := r.ParseForm()
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

            // Preenche os campos do agendamento com os valores do formulário
            var agendamento models.Agendamento
            agendamento.Nome = r.FormValue("nome_morador")
            agendamento.Apartamento = r.FormValue("apartamento")
            agendamento.Bloco = r.FormValue("bloco")
            agendamento.Local = r.FormValue("local")
            agendamento.Dia, _ = strconv.Atoi(r.FormValue("dia"))
            agendamento.Mes, _ = strconv.Atoi(r.FormValue("mes"))
            agendamento.Ano, _ = strconv.Atoi(r.FormValue("ano"))
            agendamento.Periodo = r.FormValue("periodo")
            agendamento.Funcionario = r.FormValue("funcionario")
            agendamento.Observacoes = r.FormValue("observacoes")
            agendamento.Convidados = r.FormValue("convidados")

            log.Printf("Agendamento recebido: %+v", agendamento)

            // Tenta criar o novo agendamento no banco de dados
            err = agendamento.Create(db)
            if err != nil {
                log.Printf("Erro ao criar agendamento: %s", err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Redireciona para a página de agendamentos após a criação bem-sucedida
            http.Redirect(w, r, "/agendamentos", http.StatusSeeOther)
        } else {
            // Retorna um erro 405 para métodos não permitidos
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        }
    }
}