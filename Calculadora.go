package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

// Numero estrutura para representar um número inteiro recebido como JSON.
type Numero struct {
	Numero int `json:"numero"`
}

// Total estrutura para manter o total atual e o histórico de operações.
type Total struct {
	Total     float64 `json:"total"`     // Total atual da calculadora
	historico []int   `json:"historico"` // Histórico de números utilizados nas operações
}

var num Numero  // Variável global para armazenar o número recebido nas operações
var total Total // Variável global para manter o estado total da calculadora

// adicao é um handler para a rota POST /adicao que adiciona o número recebido ao total.
func adicao(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&num) // Decodifica o corpo da requisição JSON para num
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Retorna erro 400 se houver erro na decodificação
		return
	}
	total.historico = append(total.historico, num.Numero) // Adiciona o número ao histórico
	total.Total += float64(num.Numero)                    // Realiza a operação de adição no total

	response := struct {
		Numero Numero `json:"Num"`   // Resposta JSON contendo o número recebido
		Total  int    `json:"Total"` // Total atual convertido para inteiro
	}{
		Numero: num,
		Total:  int(total.Total),
	}

	err = json.NewEncoder(w).Encode(response) // Codifica a resposta para JSON e envia de volta ao cliente
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Retorna erro 500 se houver erro na codificação
		return
	}
}

// subtracao é um handler para a rota POST /subtracao que subtrai o número recebido do total.
func subtracao(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	total.historico = append(total.historico, num.Numero)
	total.Total -= float64(num.Numero)

	response := struct {
		Numero Numero `json:"Num"`
		Total  int    `json:"Total"`
	}{
		Numero: num,
		Total:  int(total.Total),
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// multiplicacao é um handler para a rota POST /multiplicacao que multiplica o total pelo número recebido.
func multiplicacao(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	total.historico = append(total.historico, num.Numero)
	total.Total *= float64(num.Numero)

	response := struct {
		Numero Numero `json:"Num"`
		Total  int    `json:"Total"`
	}{
		Numero: num,
		Total:  int(total.Total),
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// divisao é um handler para a rota POST /divisao que divide o total pelo número recebido.
func divisao(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	total.historico = append(total.historico, num.Numero)

	// Verifica divisão por zero
	if num.Numero == 0 {
		http.Error(w, "Divisão por 0 não permitida", http.StatusBadRequest)
		return
	}

	total.Total /= float64(num.Numero)

	response := struct {
		Numero Numero `json:"Num"`
		Total  int    `json:"Total"`
	}{
		Numero: num,
		Total:  int(total.Total),
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// raizquadrada é um handler para a rota POST /raizquadrada que calcula a raiz quadrada do número recebido.
func raizquadrada(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//verificar se o numero não e negativo
	if num.Numero < 0 {
		http.Error(w, "Nao e possivel calcular raiz de 0", http.StatusBadRequest)
		return
	}
	raiz := math.Sqrt(float64(num.Numero))
	total.historico = append(total.historico, num.Numero)
	response := struct {
		Numero Numero  `json:"Num"`
		Raiz   float64 `json:"Raiz"`
	}{
		Numero: num,
		Raiz:   raiz,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// delete é um handler para a rota DELETE /delete que reinicia o total da calculadora.
func delete(w http.ResponseWriter, r *http.Request) {
	total.Total = 0                       // Reinicia o total para zero
	fmt.Fprintf(w, "Limpando o total...") // Informa ao cliente que o total foi reiniciado
	w.WriteHeader(http.StatusNoContent)   // Retorna status 204 - No Content
}

// seetotal é um handler para a rota GET /seetotal que retorna o valor atual do total da calculadora.
func seetotal(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&total.Total) // Codifica o total atual para JSON e envia ao cliente
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Retorna erro 500 se houver erro na codificação
		return
	}
}

// homepage é um handler para a rota GET /homepage que fornece uma breve descrição da API.
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Calculadora.\n Adiciocar numeros no POST /adicao, subtracao, multiplicacao, divisao\n -Ver total GET /total\n Para limpar a conta /delete ")
}

func main() {
	r := mux.NewRouter() // Inicializa o roteador Gorilla Mux

	log.Println("Iniciando API calculadora... ")

	// Definição das rotas da API com os handlers correspondentes
	r.HandleFunc("/homepage", homepage).Methods("GET")
	r.HandleFunc("/adicao", adicao).Methods("POST")
	r.HandleFunc("/subtracao", subtracao).Methods("POST")
	r.HandleFunc("/multiplicacao", multiplicacao).Methods("POST")
	r.HandleFunc("/divisao", divisao).Methods("POST")
	r.HandleFunc("/raizquadrada", raizquadrada)
	r.HandleFunc("/delete", delete).Methods("DELETE")
	r.HandleFunc("/seetotal", seetotal).Methods("GET")

	// Inicia o servidor HTTP na porta 8080 com o roteador configurado
	log.Fatal(http.ListenAndServe(":8080", r))
}
