package count

import (
	"fmt"
	"strings"
)

func CountWord(word string) map[string]int {
	 mapWord := map[string]int{}
	// Fields sirve para separar palabras, pero no deja espacios vacios y toma saltos de linea, es mejor que split en separar palabras.
	splitWords := strings.Fields(word)
	for _,value:= range (splitWords) {
			mapWord[value]+= 1

		// Como comprobamos si la clave existe
		// if count, existe := mapWord[word]; existe {
            // *Ya existe, es duplicada
        //     mapWord[word] = count + 1
        // } else {
            // *Primera vez que aparece
        //     mapWord[word] = 1
        // }


	}

	fmt.Println(mapWord)

	return mapWord
}