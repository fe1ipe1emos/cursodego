package manipulador

import (
	"fmt"
	"net/http"
)

//Funcao é o que vai ver o request do /funcao
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipulador usando função em um pacote!")
}
