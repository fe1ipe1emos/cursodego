package manipulador

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fe1ipe1emos/cursodego/avancado/banco_sql/model"
	"github.com/fe1ipe1emos/cursodego/avancado/banco_sql/repo"
)

//Local é o manipulador da requisição /local
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}

	//Query será definida pelo endereço da página!
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Numero inválido, verifique!", http.StatusBadRequest)
		fmt.Println("[local] Erro ao converter o numero enviado ", err.Error())
	}

	sqlQuery := "SELECT country, city, telcode FROM cursodego.place WHERE telcode = ?"
	linhas, err := repo.Db.Queryx(sqlQuery, codigoTelefone)
	if err != nil {
		http.Error(w, "Nao foi possivel pesquisar este numero ", http.StatusInternalServerError)
		fmt.Println("[local] Erro ao converter o numero enviado ", sqlQuery, "Erro ", err.Error())
		return
	}
	for linhas.Next() {
		err = linhas.StructScan(&local)
		if err != nil {
			http.Error(w, "Nao foi possivel pesquisar este numero ", http.StatusInternalServerError)
			fmt.Println("[local] Erro ao converter o numero enviado, ", sqlQuery, "Erro ", err.Error())
			return
		}
	}

	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "Deu ruim na renderização da página", http.StatusInternalServerError)
		fmt.Println("Erro na execução do modelo: ", err.Error())
	}

	/*
		Registrar navegação do usuário em um site (exemplo de persistencia)
	*/

	sqlInsert := "INSERT INTO cursodego.logquery (daterequest) values (?)"
	resultado, err := repo.Db.Exec(sqlInsert, time.Now().Format("02/01/2006 15:04:05"))
	if err != nil {
		fmt.Println("[local] Erro ao inserir o dado, ", sqlInsert, "Erro ", err.Error())
	}
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		fmt.Println("[local] Erro ao pegar o numero de linhas ", sqlInsert, "Erro ", err.Error())
	}
	fmt.Println("Linhas afetadas: ", linhasAfetadas)

}
