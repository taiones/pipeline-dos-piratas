package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// Testa se a rota /nakamas responde corretamente
func TestNakamasHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/nakamas", nil) // cria uma requisição falsa
    w := httptest.NewRecorder()                        // grava a resposta

    nakamasHandler(w, req) // chama a função que criamos no main.go

    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Esperava status 200, mas veio %d", resp.StatusCode)
    }
}
