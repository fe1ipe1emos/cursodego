package model

import "database/sql"

//Local armazena (reflete) os dados da localidade de acordo com o código telefonico.
type Local struct {
	//Etiqueta (tag) db para SQL, de acordo com o nome no banco!
	//Cidade está com sql.NullString pois no banco ele não está como not null, logo esse pacote evita campos nulos!

	Pais             string         `json:"pais" db:"country"`
	Cidade           sql.NullString `json:"cidade" db:"city"`
	CodigoTelefonico int            `json:"cod_telefone" db:"telcode"`
}
