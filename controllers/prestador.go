package controllers

import (
    "condominio/models"
    "database/sql"
    "html/template"
    "log"
    "net/http"
)

// PrestadorFormHandler lida com as requisições HTTP para exibir o formulário de prestadores.
// Suporta apenas o método GET para renderizar a página de prestadores.
func PrestadorFormHandler(w http.ResponseWriter, r *http.Request) {
    // Carrega e renderiza o template prestadores.html para requisições GET.
    tmpl, err := template.ParseFiles("templates/prestadores.html")
    if err != nil {
        // Retorna um erro 500 se ocorrer um problema ao carregar o template.
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

// PrestadorHandler lida com as requisições HTTP para criar novos prestadores.
// Suporta apenas o método POST para criar um novo prestador.
func PrestadorHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            // Trata requisições POST para criar um novo prestador.
            err := r.ParseForm()
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

            // Preenche os campos do prestador com os valores do formulário.
            prestador := models.Prestador{
                NomeEmpresa:     r.FormValue("nome-empresa"),
                TipoServico:     r.FormValue("tipo-servico"),
                Data:            r.FormValue("data"),
                HoraEntrada:     r.FormValue("hora-entrada"),
                NomePrestador:   r.FormValue("nome-prestador"),
                RGCPF:           r.FormValue("rg-cpf"),
                Telefone:        r.FormValue("telefone"),
                ContratanteNome: r.FormValue("contratante-nome"),
                ContratanteTipo: r.FormValue("contratante-tipo"),
                ContratanteApto: r.FormValue("contratante-apto"),
                ContratanteBloco: r.FormValue("contratante-bloco"),
                Autorizou:       r.FormValue("autorizou"),
            }

            horaSaida := r.FormValue("hora-saida")
            if horaSaida != "" {
                prestador.HoraSaida = &horaSaida
            } else {
                prestador.HoraSaida = nil
            }

            log.Printf("Prestador recebido: %+v", prestador)

            // Tenta criar o novo prestador no banco de dados.
            err = prestador.Create(db)
            if err != nil {
                log.Printf("Erro ao criar prestador: %s", err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Redireciona para a página de sucesso após a criação bem-sucedida.
            http.Redirect(w, r, "/prestadores", http.StatusSeeOther)
        } else {
            // Redireciona para a página de sucesso se o método não for POST.
            http.Redirect(w, r, "/prestadores", http.StatusFound)
        }
    }
}