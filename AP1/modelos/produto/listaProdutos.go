package produto

var listaProdutos []Produto

func GetListaProdutos() *[]Produto { return &listaProdutos }

func ListaVazia() bool { return len(listaProdutos) == 0 }

// Retorta o produto de index i de listaProdutos
func GetProduto(i int) *Produto {
	if i < 0 || i > len(listaProdutos) {
		return nil
	}
	return &listaProdutos[i]
}

// Adiciona um produto no final de listaProdutos
func AdicionarProudto(tmp *Produto) {
	listaProdutos = append(listaProdutos, *tmp)
}

// Retorna o index de um produto baseado no id
func AcharProduto(id int) int {
	for i, produto := range listaProdutos {
		if produto.ID == id {
			return i
		}
	}
	return -1
}

// Remove o produto de index i de listaProdutos
func RemoverProduto(i int) {
	if i < 0 || i >= len(listaProdutos) {
		return
	}
	listaProdutos = append(listaProdutos[:i], listaProdutos[i+1:]...)
}
