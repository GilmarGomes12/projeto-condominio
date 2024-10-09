package controllers

import (
    "encoding/json"
    "net/http"
    "condominio/models"
    "database/sql"
)

// PesquisaAgendamentoHandler lida com as requisições HTTP para pesquisar agendamentos.
// Suporta apenas o método GET para buscar agendamentos com base em um parâmetro de consulta.
func PesquisaAgendamentoHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Obtém o parâmetro de consulta "query" da URL.
        query := r.URL.Query().Get("query")
        if query == "" {
            // Retorna um erro 400 se o parâmetro de consulta estiver ausente.
            http.Error(w, "Query parameter is required", http.StatusBadRequest)
            return
        }

        // Busca os agendamentos no banco de dados com base na query.
        agendamentos, err := (&models.Agendamento{}).Search(db, query)
        if err != nil {
            // Retorna um erro 500 se ocorrer um problema na busca.
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Define o cabeçalho da resposta como JSON e retorna os agendamentos encontrados.
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(agendamentos)
    }
}