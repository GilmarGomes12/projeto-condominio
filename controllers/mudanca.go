package controllers

import (
    "condominio/models"
    "database/sql"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "time"
)

// MudancaHandler lida com as requisições HTTP para agendamento de mudanças.
// Suporta métodos POST para criar um novo agendamento e GET para renderizar a página de agendamento.
func MudancaHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            // Trata requisições POST para criar um novo agendamento de mudança.
            var agendamento models.Mudanca

            // Converte a data da mudança do formato string para time.Time.
            dataMudanca, err := time.Parse("2006-01-02", r.FormValue("data-mudanca"))
            if err != nil {
                http.Error(w, "Data inválida", http.StatusBadRequest)
                return
            }
            agendamento.DataMudanca = dataMudanca

            agendamento.ResponsavelNome = r.FormValue("responsavel-nome")
            agendamento.ResponsavelApto = r.FormValue("responsavel-apto")
            agendamento.ResponsavelBloco = r.FormValue("responsavel-bloco")
            agendamento.Horario = r.FormValue("horario")
            agendamento.NomeEmpresa = r.FormValue("nome-empresa")

            usoElevador, err := strconv.ParseBool(r.FormValue("uso-elevador"))
            if err != nil {
                usoElevador = false
            }
            agendamento.UsoElevador = usoElevador

            usoEscada, err := strconv.ParseBool(r.FormValue("uso-escada"))
            if err != nil {
                usoEscada = false
            }
            agendamento.UsoEscada = usoEscada

            iscar, err := strconv.ParseBool(r.FormValue("iscar"))
            if err != nil {
                iscar = false
            }
            agendamento.Iscar = iscar

            // Tenta criar o novo agendamento no banco de dados.
            err = agendamento.Create(db)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Redireciona o usuário de volta para a página do formulário.
            http.Redirect(w, r, "/mudanca", http.StatusSeeOther)
        case "GET":
            // Trata requisições GET para renderizar a página de agendamento de mudança.
            tmpl, err := template.ParseFiles("templates/mudanca.html")
            if err != nil {
                log.Printf("Erro ao carregar o template: %v", err)
                http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
                return
            }
            tmpl.Execute(w, nil)
        default:
            // Retorna um erro 405 para métodos não permitidos.
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        }
    }
}