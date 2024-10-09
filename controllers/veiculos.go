package controllers

import (
    "html/template"
    "net/http"
    "log"
)

// VeiculosHandler lida com as requisições HTTP para exibir a página de veículos.
// Suporta apenas o método GET para renderizar a página de veículos.
func VeiculosHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Carrega e renderiza o template veiculos.html para requisições GET.
        tmpl, err := template.ParseFiles("templates/veiculos.html")
        if err != nil {
            // Registra o erro e retorna um erro 500 se ocorrer um problema ao carregar o template.
            log.Printf("Erro ao carregar o template: %v", err)
            http.Error(w, "Erro ao carregar o template", http.StatusInternalServerError)
            return
        }
        tmpl.Execute(w, nil)
    }
}