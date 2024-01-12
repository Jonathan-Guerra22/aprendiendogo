package main

import (
	"GODESDE0/variables"
	"fmt"
)

func main(){
	//variables.MuestroEnteros()	
	//variables.RestoVariables()
	estado, texto := variables.ConvertirATexto(21312)

	fmt.Println("Estado = ", estado)
	fmt.Println("Texto = ", texto)

}