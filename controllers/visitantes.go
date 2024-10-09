package controllers

import (
    "condominio/models"
    "database/sql"
    "html/template"
    "log"
    "net/http"

)

// VisitantesFormHandler lida com as requisições HTTP para exibir o formulário de visitantes.
// Suporta apenas o método GET para renderizar a página de visitantes.
func VisitantesFormHandler(w http.ResponseWriter, r *http.Request) {
    // Carrega e renderiza o template visitantes.html para requisições GET.
    tmpl, err := template.ParseFiles("templates/visitantes.html")
    if err != nil {
        // Retorna um erro 500 se ocorrer um problema ao carregar o template.
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

// VisitantesHandler lida com as requisições HTTP para criar novos visitantes.
// Suporta apenas o método POST para criar um novo visitante.
func VisitantesHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            // Parseia o formulário
            err := r.ParseForm()
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }


             // Preenche os campos do visitante com os valores do formulário
            var visitante models.Visitante
            visitante.Data = r.FormValue("data")
            visitante.Nome = r.FormValue("nome-visitante")
            visitante.Documento = r.FormValue("rg-cpf")
            visitante.Visitando = r.FormValue("visitando")
            visitante.Apartamento = r.FormValue("apartamento")
            visitante.Bloco = r.FormValue("bloco")
            visitante.HoraEntrada = r.FormValue("hora-entrada")
            
            horaSaida := r.FormValue("hora-saida")
            if horaSaida != "" {
                visitante.HoraSaida = &horaSaida
            } else {
                visitante.HoraSaida = nil
            }

            visitante.Autorizou = r.FormValue("autorizou")
            visitante.Vaga = r.FormValue("vaga")
            visitante.Placa = r.FormValue("placa")
            visitante.Marca = r.FormValue("marca")
            visitante.Modelo = r.FormValue("modelo")
            visitante.Cor = r.FormValue("cor")

            log.Printf("Visitante recebido: %+v", visitante)


            // Tenta criar o novo visitante no banco de dados
            err = visitante.Create(db)
            if err != nil {
                log.Printf("Erro ao criar visitante: %s", err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Redireciona para a página de visitantes após a criação bem-sucedida
            http.Redirect(w, r, "/visitantes", http.StatusSeeOther)
        } else {
            // Retorna um erro 405 para métodos não permitidos
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        }
    }
}