package pedidos

import (
	"AP1/handlers/metricas"
	"AP1/modelos/pedido"
	"AP1/modelos/produto"
	"encoding/json"
	"fmt"
	"net/http"
)

var idPedido int = 1

func IncluirPedido(w http.ResponseWriter, r *http.Request) {
	var tmp pedido.Pedido
	//Decodificação do json
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Criar mensagem de resposta
	resposta := fmt.Sprintf("Pedido criado com sucesso (ID = %d)", idPedido)
	if tmp.ID != 0 {
		resposta += "\n[Aviso: ID fornecido no json desconsiderado]"
	}
	if tmp.ValorTotal != 0 {
		resposta += "\n[Aviso: Valor fornecido no json desconsiderado]"
	}

	tmp.ValorTotal = 0

	//Verifica se todos os produtos do pedido são validos
	for _, i := range tmp.Produtos {
		index := produto.AcharProduto(i.ProdutoID)
		if index == -1 {
			msg := fmt.Sprintf("produto com ID %d não existe", i.ProdutoID)
			http.Error(w, msg, http.StatusBadRequest)
			return
		} else {
			produtoTmp := produto.GetProduto(index)
			//Incrementa o valor total
			tmp.ValorTotal += produtoTmp.Valor * float32(i.Quantidade)
		}
	}

	tmp.ID = idPedido

	//Inclue a taxa de entrega
	if tmp.Delivery {
		tmp.ValorTotal += 10
	}

	//Adiciona o pedido a fila
	pedido.IncluirPedido(&tmp)

	//Resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, resposta)

	idPedido++ //Incrementar o contador de id para o próximo pedido

	metricas.IncPedidosEmAndamento(1)
	metricas.SomFaturamentoTotal(tmp.ValorTotal)
}

func ExpedirPedido() bool {
	if pedido.ExpedirPedido() {
		metricas.SomPedidosEncerrados()
		return true
	}
	return false
}

func ListarPedidos(w http.ResponseWriter, r *http.Request) {
	if pedido.FilaVazia() {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Não ha pedidos em andamento")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*(pedido.GetFilaPedidos()))
}
