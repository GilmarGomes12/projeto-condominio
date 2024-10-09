package controllers

import (
    "encoding/json"
    "html/template"
    "net/http"
    "log"
    "io/ioutil"
    "condominio/models"
    "database/sql"
    "time"
    "github.com/google/uuid"
)

// EncomendasHandler lida com requisições HTTP para o recurso de encomendas.
// Suporta métodos POST para criar uma nova encomenda e GET para buscar encomendas.
func EncomendasHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            // Lê o corpo da requisição.
            body, err := ioutil.ReadAll(r.Body)
            if err != nil {
                log.Printf("Erro ao ler o corpo da requisição: %v", err)
                http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
                return
            }

            log.Printf("Corpo da requisição recebido: %s", string(body))

            var encomenda models.Encomenda
            // Decodifica o JSON recebido no corpo da requisição.
            err = json.Unmarshal(body, &encomenda)
            if err != nil {
                log.Printf("Erro ao decodificar JSON: %v", err)
                http.Error(w, "Erro ao decodificar JSON: " + err.Error(), http.StatusBadRequest)
                return
            }

            // Gera um número de protocolo único e define a data/hora de recebimento.
            encomenda.NumeroProtocolo = uuid.New().String()
            encomenda.DataHoraRecebimento = time.Now()

            // Cria a nova encomenda no banco de dados.
            err = encomenda.Create(db)
            if err != nil {
                log.Printf("Erro ao criar encomenda: %v", err)
                http.Error(w, "Erro ao criar encomenda", http.StatusInternalServerError)
                return
            }

            // Retorna a encomenda criada com status 201 Created.
            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(encomenda)
        case "GET":
            // Obtém o parâmetro de consulta "query" da URL.
            query := r.URL.Query().Get("query")
            if query == "" {
                // Se não houver query, carrega e exibe o template HTML.
                tmpl, err := template.ParseFiles("templates/encomendas.html")
                if err != nil {
                    log.Printf("Erro ao carregar o template: %v", err)
                    http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
                    return
                }
                tmpl.Execute(w, nil)
                return
            }

            var (
                encomendas []models.Encomenda
                err        error
            )
            // Busca encomendas no banco de dados com base na query.
            if query == "" {
                encomendas, err = (&models.Encomenda{}).GetAll(db)
            } else {
                encomendas, err = (&models.Encomenda{}).Search(db, query)
            }
            if err != nil {
                log.Printf("Erro ao buscar encomendas: %v", err)
                http.Error(w, "Erro ao buscar encomendas", http.StatusInternalServerError)
                return
            }

            // Retorna as encomendas encontradas como JSON.
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(encomendas)
        default:
            // Retorna erro 405 Method Not Allowed para métodos não suportados.
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        }
    }
}