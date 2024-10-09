package controllers

import (
    "encoding/json"
    "html/template"
    "net/http"
    "log"
    "condominio/models"
    "database/sql"
)

// ConfiguracoesHandler lida com as requisições HTTP para a página de configurações.
// Suporta métodos POST para criar uma nova configuração e GET para renderizar a página de configurações.
func ConfiguracoesHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            // Trata requisições POST para criar uma nova configuração.
            var configuracao models.Configuracao
            err := json.NewDecoder(r.Body).Decode(&configuracao)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

            // Tenta criar a nova configuração no banco de dados.
            err = configuracao.Create(db)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Retorna a configuração criada como resposta JSON.
            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(configuracao)
        case "GET":
            // Trata requisições GET para renderizar a página de configurações.
            tmpl, err := template.ParseFiles("templates/configuracoes.html")
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