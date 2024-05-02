package pedido

var filaPedidos []Pedido

func GetFilaPedidos() *[]Pedido { return &filaPedidos }

func FilaVazia() bool { return len(filaPedidos) == 0 }

func IncluirPedido(tmp *Pedido) {
	filaPedidos = append(filaPedidos, *tmp)
}

func ExpedirPedido() bool {
	if len(filaPedidos) > 0 {
		filaPedidos = filaPedidos[1:]
		return true
	}
	return false
}
