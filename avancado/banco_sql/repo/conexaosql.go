package repo

import (
	/*
		github.com/go-sql-driver/mysql não será usado pela aplicação
	*/
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
Neste arquivo estará a abertura da conexão com o banco de dados
 Importante: O ideal é não ficar abrindo e fechando conexão com o banco de dados.
             Deixar a conexao publica!
*/

//Db aramazena a conexao  com o banco de dados
var Db *sqlx.DB

//AbreConexaoDB é a função para abrir conexão com o banco de dados!
func AbreConexaoDB() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root:pazuzinho@tcp(localhost:3306)/cursodego")
	if err != nil {
		return
	}
	//precisa pingar o banco pra ver se está tudo ok!
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}
