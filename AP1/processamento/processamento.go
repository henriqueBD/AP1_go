package processamento

import (
	"AP1/handlers/pedidos"
	"fmt"
	"time"
)

func Processar(lojaPonteiro *bool) {
	for {
		time.Sleep(30 * time.Second)
		if *lojaPonteiro {
			if pedidos.ExpedirPedido() {
				fmt.Println("Pedido expedido")
			}
		} else {
			return
		}
	}
}
