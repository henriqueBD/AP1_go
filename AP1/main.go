package main

import (
	"AP1/handlers/loja"
	"AP1/handlers/metricas"
	"AP1/handlers/pedidos"
	"AP1/handlers/produtos"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/abrir", loja.AbrirLoja).Methods("POST")
	r.HandleFunc("/fechar", loja.FecharLoja).Methods("POST")

	r.HandleFunc("/produto", produtos.CriarProduto).Methods("POST")
	r.HandleFunc("/produto", produtos.ListarProduto).Methods("GET")
	r.HandleFunc("/produto", produtos.RemoverProduto).Methods("DELETE")
	r.HandleFunc("/produtos", produtos.ListarProdutos).Methods("GET")

	r.HandleFunc("/pedido", pedidos.IncluirPedido).Methods("POST")
	r.HandleFunc("/pedidos", pedidos.ListarPedidos).Methods("GET")

	r.HandleFunc("/metricas", metricas.GetMetricas).Methods("GET")

	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
