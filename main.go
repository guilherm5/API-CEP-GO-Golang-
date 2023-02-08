package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CEP struct {
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	Estado      string `json:"uf"`
}

func main() {

	// variavel que vai ser printada com os dados do cep requisitado
	var dadosCEP CEP
	//string que vai receber meu cep
	var cep string
	//pausando meu programa para digitar meu cep
	fmt.Scanf("%s", &cep)
	// puxando informações (consumindo api via cep) ("+ cep +" esta esperando digitar o cep)
	resposta, _ := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	// ioutill para ler toda minha requisição
	dados, _ := ioutil.ReadAll(resposta.Body)
	defer resposta.Body.Close()

	//json unmarshal transforma um json em um struct ou map etc..
	json.Unmarshal(dados, &dadosCEP)

	//print do resultado
	fmt.Println(dadosCEP)

}
