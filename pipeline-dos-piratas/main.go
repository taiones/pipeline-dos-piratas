package main

import (
    "fmt"
    "net/http"
)

// Função que responde quando alguém acessa /nakamas
func nakamasHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, `["Luffy","Zoro","Nami","Sanji"]`)
}

// Função principal que inicia o servidor
func main() {
    http.HandleFunc("/nakamas", nakamasHandler) // liga a rota /nakamas à função acima
    fmt.Println("Servidor rodando na porta 8080...")
    http.ListenAndServe(":8080", nil) // inicia o servidor na porta 8080
}
