package controllers

import (
    "encoding/json"
    "html/template"
    "net/http"
    "log"
    "condominio/models"
    "database/sql"
    "time"
)

// OcorrenciaHandler lida com as requisições HTTP para gerenciar ocorrências.
func OcorrenciaHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            var ocorrencia models.Ocorrencia
            err := json.NewDecoder(r.Body).Decode(&ocorrencia)
            if err != nil {
                log.Printf("Erro ao decodificar JSON: %v", err)
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

            // Tenta converter a data de ocorrência do formato string para time.Time.
            ocorrencia.DataOcorrencia, err = time.Parse(time.RFC3339, ocorrencia.DataOcorrencia.Format(time.RFC3339))
            if err != nil {
                log.Printf("Erro ao analisar a data de ocorrência: %v", err)
                http.Error(w, "Formato de data inválido", http.StatusBadRequest)
                return
            }
            ocorrencia.DataRegistro, err = time.Parse(time.RFC3339, ocorrencia.DataRegistro.Format(time.RFC3339))
            if err != nil {
                log.Printf("Data de registro inválida: %v", err)
                http.Error(w, "Data de registro inválida", http.StatusBadRequest)
                return
            }

            // Tenta criar a nova ocorrência no banco de dados.
            err = ocorrencia.Create(db)
            if err != nil {
                log.Printf("Erro ao criar ocorrência: %v", err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Retorna a ocorrência criada como resposta JSON.
            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(ocorrencia)
        case "GET":
            // Trata requisições GET para listar ou buscar ocorrências.
            query := r.URL.Query().Get("query")
            if query == "" {
                // Renderiza a página de ocorrências se não houver query.
                tmpl, err := template.ParseFiles("templates/ocorrencias.html")
                if err != nil {
                    log.Printf("Erro ao carregar o template: %v", err)
                    http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
                    return
                }
                tmpl.Execute(w, nil)
                return
            }

            // Busca ocorrências no banco de dados.
            ocorrencias, err := (&models.Ocorrencia{}).Search(db, query)
            if err != nil {
                log.Printf("Erro ao buscar ocorrências: %v", err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Retorna as ocorrências encontradas como resposta JSON.
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(ocorrencias)
        default:
            // Retorna um erro 405 para métodos não permitidos.
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        }
    }
}