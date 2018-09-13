package main

import (
	"fmt"
	"net/http"

	"github.com/fe1ipe1emos/cursodego/avancado/banco_sql/manipulador"
	"github.com/fe1ipe1emos/cursodego/avancado/banco_sql/repo"
)

func main() {

	err := repo.AbreConexaoDB()
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados! ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bem vindo a página inicial!")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)

	fmt.Println("Estou on!")
	http.ListenAndServe(":8181", nil)
}
