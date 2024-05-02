package produtos

import (
	"AP1/handlers/metricas"
	"AP1/modelos/produto"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var idProduto int = 1

func CriarProduto(w http.ResponseWriter, r *http.Request) {
	var tmp produto.Produto
	//Decodificação do json
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err != nil {
		http.Error(w, "Erro em decodificar json", http.StatusBadRequest)
		return
	}

	//Verificar se o json tem todas as informações
	if tmp.Descricao == "" || tmp.Nome == "" || tmp.Valor == 0 {
		http.Error(w, "Json sem informações necessárias", http.StatusBadRequest)
		return
	}

	//Criar mensagem de resposta
	resposta := fmt.Sprintf("Produto criado com sucesso (ID = %d)", idProduto)
	if tmp.ID != 0 {
		resposta += "\n[Aviso: ID fornecido no json desconsiderado]"
	}

	tmp.ID = idProduto //Adicionar id único para o produto
	produto.AdicionarProudto(&tmp)

	//Mandar resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, resposta)

	idProduto++ //Incrementar o contador de id para o próximo produto
	metricas.IncProdutosCadastrados(1)
}

func ListarProduto(w http.ResponseWriter, r *http.Request) {
	//Conversão de string para int
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	//Caso o id seja valido, continuar o código
	index := produto.AcharProduto(id)
	if index == -1 {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	//Mandar os dados do produto
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*produto.GetProduto(index))
}

func RemoverProduto(w http.ResponseWriter, r *http.Request) {
	//Conversão de string para int
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	//Caso o id seja valido, remover o produto
	i := produto.AcharProduto(id)
	if i == -1 {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}
	produto.RemoverProduto(i)
	metricas.IncProdutosCadastrados(-1)

	//Mandar respota
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Produto removido com sucesso"))
}

func ListarProdutos(w http.ResponseWriter, r *http.Request) {
	if produto.ListaVazia() {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Não ha produtos cadastrados")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*(produto.GetListaProdutos()))
}
