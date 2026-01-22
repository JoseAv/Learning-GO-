package procesador

import (
	"slices"
	"strings"
)


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
	return strings.Join(parts, "")
}

func MultiProcesador(data string, procesadores ...Procesador)string {
	result:=data
	for _,p := range procesadores{
		result= p.Procesar(result)
	}
	return result

}
