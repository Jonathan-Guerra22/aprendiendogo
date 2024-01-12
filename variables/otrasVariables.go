package variables

import (
	"fmt"
	"time"
	"strconv"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables (){
	
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1232.212
	Fecha = time.Now()

	fmt.Println("Nombre = ", Nombre)
	fmt.Println("Estado = ", Estado)
	fmt.Println("Sueldo = ", Sueldo)
	fmt.Println("Fecha = ", Fecha)

}


func ConvertirATexto(numero int) (bool, string) {
	return true, strconv.Itoa(numero)
}