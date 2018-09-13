package main

import (
	"fmt"
	"net/http"

	"github.com/fe1ipe1emos/cursodego/avancado/servidor_web/manipulador"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bem vindo a página inicial!")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)

	fmt.Println("Estou on!")
	http.ListenAndServe(":8181", nil)
}
