package metrica

type MetricasSistema struct {
	ProdutosCadastrados int     `json:"produtosCadastrados"`
	PedidosEncerrados   int     `json:"pedidosEncerrados"`
	PedidosEmAndamento  int     `json:"pedidosEmAndamento"`
	FaturamentoTotal    float32 `json:"faturamentoTotal"`
}
