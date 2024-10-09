package controllers

import (
    "net/http"
    "html/template"
    "log"
)

// MenuHandler lida com as requisições HTTP para a página de menu.
// Suporta apenas o método GET para renderizar a página de menu.
func MenuHandler(w http.ResponseWriter, r *http.Request) {
    // Carrega e renderiza o template menu.html para requisições GET.
    tmpl, err := template.ParseFiles("templates/menu.html")
    if err != nil {
        log.Fatalf("Erro ao carregar o template: %v", err)
    }
    tmpl.Execute(w, nil)
}
