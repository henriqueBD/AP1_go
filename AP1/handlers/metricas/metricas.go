package metricas

import (
	"AP1/modelos/metrica"
	"encoding/json"
	"net/http"
)

var metricas metrica.MetricasSistema

func GetMetricas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metricas)
}

func IncProdutosCadastrados(i int)  { metricas.ProdutosCadastrados += i }
func IncPedidosEmAndamento(i int)   { metricas.PedidosEmAndamento += i }
func SomFaturamentoTotal(i float32) { metricas.FaturamentoTotal += i }

func SomPedidosEncerrados() {
	metricas.PedidosEncerrados++
	metricas.PedidosEmAndamento--
}
