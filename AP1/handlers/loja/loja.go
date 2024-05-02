package loja

import (
	"AP1/processamento"
	"fmt"
	"net/http"
)

var lojaAberta *bool

func AbrirLoja(w http.ResponseWriter, r *http.Request) {
	if lojaAberta != nil && *lojaAberta == true {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Loja ja foi aberta anteriormente")
		return
	}

	lojaAberta = new(bool)
	*lojaAberta = true
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Loja aberta com sucesso")
	go processamento.Processar(lojaAberta)
}

func FecharLoja(w http.ResponseWriter, r *http.Request) {
	if lojaAberta == nil {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Loja nao foi aberta")
	} else if *lojaAberta == false {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Loja ja fechada")
	} else {
		*lojaAberta = false
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Loja fechada com sucesso")
	}
}
