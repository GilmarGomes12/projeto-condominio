package controllers

import (
    "net/http"
)

// IndexHandler lida com a rota principal
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Bem-vindo à página inicial!"))
}
