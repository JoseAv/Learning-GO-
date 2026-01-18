package main

import (
	"fmt"
	"process/procesador"
)

func main() {

	mayus := procesador.MayusculasProcesador{}
	fmt.Println(mayus.Procesar("mayusculas"))
	inver := procesador.InvertirProcesador{}
	fmt.Println(inver.Procesar("invertido"))
	limp := procesador.LimpiarEspaciosProcesador{}
	fmt.Println(limp.Procesar("   limpio    "))

	fmt.Println(procesador.MultiProcesador("  Data Procesada  ",limp,mayus,inver))
}