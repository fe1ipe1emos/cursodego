package manipulador

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fe1ipe1emos/cursodego/avancado/banco_sql/model"
)

//Ola é uma funcao que vai fazer o handler do reqest do /ola
func Ola(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := ModeloOla.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Deu ruim na renderização da página", http.StatusInternalServerError)
		fmt.Println("Erro na execução do modelo: ", err.Error())
	}
}
