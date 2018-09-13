package manipulador

import "html/template"

//ModeloOla é o modelo /ola
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal é o modelo /local
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
