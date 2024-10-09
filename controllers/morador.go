package controllers

import (
    "condominio/models"
    "database/sql"
    "log"
    "net/http"
    "html/template"
)

// TemplateData é uma estrutura que pode ser usada para passar dados para os templates.
type TemplateData struct {
    Success bool
}

// MoradorFormHandler lida com as requisições HTTP para exibir o formulário de moradores.
// Suporta apenas o método GET para renderizar a página de moradores.
func MoradorFormHandler(w http.ResponseWriter, r *http.Request) {
    // Carrega e renderiza o template moradores.html para requisições GET.
    tmpl, err := template.ParseFiles("templates/moradores.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

// MoradorHandler lida com as requisições HTTP para criar novos moradores.
// Suporta apenas o método POST para criar um novo morador.
func MoradorHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            // Trata requisições POST para criar um novo morador.
            var morador models.Morador
            err := r.ParseForm()
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

            // Preenche os campos do morador com os valores do formulário.
            morador.Apartamento = r.FormValue("apartamento")
            morador.Bloco = r.FormValue("bloco")
            morador.Telefone1 = r.FormValue("telefone1")
            morador.Telefone2 = r.FormValue("telefone2")
            morador.Email = r.FormValue("email")
            morador.Email2 = r.FormValue("email2")
            morador.Observacao = r.FormValue("observacao")

            // Preenche a lista de moradores com os valores do formulário.
            nomesMoradores := r.Form["nome_morador[]"]
            datasNascimento := r.Form["data_nascimento[]"]
            for i := range nomesMoradores {
                morador.Moradores = append(morador.Moradores, models.MoradorInfo{
                    Nome:           nomesMoradores[i],
                    DataNascimento: datasNascimento[i],
                })
            }

            // Preenche a lista de veículos com os valores do formulário.
            placas := r.Form["placa[]"]
            cores := r.Form["cor[]"]
            marcas := r.Form["marca[]"]
            modelos := r.Form["modelo[]"]
            for i := range placas {
                morador.Veiculos = append(morador.Veiculos, models.Veiculo{
                    Placa:  placas[i],
                    Cor:    cores[i],
                    Marca:  marcas[i],
                    Modelo: modelos[i],
                })
            }

            log.Printf("Dados recebidos: %+v\n", morador)

            // Tenta criar o novo morador no banco de dados.
            err = morador.Create(db)
            if err != nil {
                log.Printf("Erro ao criar morador: %v\n", err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Redireciona para a página de moradores após a criação bem-sucedida.
            http.Redirect(w, r, "/morador", http.StatusSeeOther)
        }
    }
}