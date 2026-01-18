package procesador

import (
	"slices"
	"strings"
)

// Ejercicio 9: Procesador de datos
// Define una interfaz Procesador con Procesar(data string) string.
// Implementa diferentes procesadores: MayusculasProcesador, InvertirProcesador,
//  LimpiarEspaciosProcesador.
//  Crea una función que aplique múltiples procesadores en cadena.


type Procesador interface{
	Procesar(data string) string
}

// Cumple con la interfaces aunque no tenga datos
type MayusculasProcesador struct{}
func (m MayusculasProcesador) Procesar(data string) string {
	return strings.ToUpper(data)
}

type LimpiarEspaciosProcesador struct{}

func (l LimpiarEspaciosProcesador) Procesar(data string)string {
	return  strings.TrimSpace(data)
}


type InvertirProcesador struct{}
func (i InvertirProcesador) Procesar(data string)string {
	parts := strings.Split(data, "")
	slices.Reverse(parts)
	return strings.Join(parts[1:], "")
}

func MultiProcesador(data string, procesadores ...Procesador)string {
	result:=data
	for _,p := range procesadores{
		result= p.Procesar(result)
	}
	return result

}
