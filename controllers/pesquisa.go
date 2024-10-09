package controllers

import (
	"condominio/models"
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// PesquisaFormHandler lida com as requisições HTTP para exibir o formulário de pesquisa.
// Suporta apenas o método GET para renderizar a página de pesquisa.
func PesquisaFormHandler(w http.ResponseWriter, r *http.Request) {
    // Carrega e renderiza o template pesquisa.html para requisições GET.
    tmpl, err := template.ParseFiles("templates/pesquisar.html")
    if err != nil {
        // Retorna um erro 500 se ocorrer um problema ao carregar o template.
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

// PesquisaHandler lida com as requisições HTTP para pesquisar diferentes tipos de dados.
// Suporta apenas o método GET para buscar dados com base em um parâmetro de consulta.
func PesquisaHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Obtém os parâmetros de consulta "query" e "tipo" da URL.
        query := r.URL.Query().Get("query")
        tipo := r.URL.Query().Get("tipo")

        if query == "" || tipo == "" {
            // Retorna um erro 400 se os parâmetros de consulta estiverem ausentes.
            http.Error(w, "Query and type parameters are required", http.StatusBadRequest)
            return
        }

        var resultados interface{}
        var err error

        // Busca os dados no banco de dados com base no tipo de pesquisa.
        switch tipo {
        case "morador":
            resultados, err = (&models.Morador{}).Search(db, query)
        case "agendamento":
            resultados, err = (&models.Agendamento{}).Search(db, query)
        case "encomenda":
            resultados, err = (&models.Encomenda{}).Search(db, query)
        case "prestador":
            resultados, err = (&models.Prestador{}).Search(db, query)
        case "mudanca":
            resultados, err = (&models.Mudanca{}).Search(db, query)
        case "visitante":
            resultados, err = (&models.Visitante{}).Search(db, query)
        case "ocorrencias":
            resultados, err = (&models.Ocorrencia{}).Search(db, query)
        case "funcionarios":
            resultados, err = (&models.Funcionarios{}).Search(db, query)
        case "domesticos":
            resultados, err = (&models.Domesticos{}).Search(db, query)                        
        default:
            http.Error(w, "Invalid search type", http.StatusBadRequest)
            return
        }

        if err != nil {
            // Retorna um erro 500 se ocorrer um problema na busca.
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Log dos resultados para depuração
        log.Printf("Resultados da pesquisa: %+v\n", resultados)

        // Define o cabeçalho da resposta como JSON e retorna os resultados encontrados.
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resultados)
    }
}