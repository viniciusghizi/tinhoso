package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Lembre-se que os pedidos tem que estar com indErro = true no banco e é necessário rodar o script no banco")

	//Nesse array PEDIDO, ficará o numero do pedido, colocar o ID do pedido
	pedido := []string{"46538650",
		"47571272",
		"48704238",
		"50805761",
		"51247231",
		"53148091"}
	//Neste array movimentacaoPedido, ficará o ID da movimentação
	movimentacaoPedido := []string{"9396233",
		"9259513",
		"9360702",
		"9247103",
		"9247015",
		"9344774"}

	file, err := os.Create("resultado.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for contador := 0; contador < len(movimentacaoPedido); contador++ {

		fmt.Println("Enviando o Pedido: ")
		fmt.Print(pedido[contador] + "\n")
		url := os.Getenv("url") + movimentacaoPedido[contador] + "?referencia=" + os.Getenv("referencia")
		method := "PUT"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			file.WriteString(err.Error())
			return
		}
		req.Header.Add("appId", "3f541b54b9ddb459bb591e486c5cf15e")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("token", os.Getenv("token"))

		res, err := client.Do(req)
		if err != nil {
			file.WriteString(err.Error())
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			file.WriteString(err.Error())
			return
		}
		str := string(body)
		var tratamento = "\n Pedido: " + pedido[contador] + " - Mensagem: " + str

		file.WriteString(tratamento)
	}
	var resultado = "\nFoi efetuado o reenvio em " + strconv.Itoa(len(pedido))
	file.WriteString(resultado)
	fmt.Println(resultado)
	fmt.Println("Após finalizar, remover o indErros de todos os pedidos e colocar nos que estão no arquivo PedidosComErro")
}
