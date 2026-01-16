package main

import (
	"fmt"
	"list/operations"
)




func main(){
 	ListWorks := []operations.Works{}
	var index int
 
	operations.AddWork("hola", &ListWorks)
	operations.AddWork("hola2", &ListWorks)
	operations.AddWork("hola3", &ListWorks)
	operations.Resume(&ListWorks)
	operations.FilterWork(false,&ListWorks)

	fmt.Println("Ingrese el index")
	// n -> numero de valores que se leyeron correctamente
	// erro -> si hubo error o no

	// no confundir con struct donde al extraer s,exist -> s si es el valor si es exitoso
	// pero si es false es el valor pode defecto de ese tipo
	
	n, err:=fmt.Scan(&index)
	fmt.Println(n)
	fmt.Println(err)

	if(err != nil){
		fmt.Println("Debe de ingresar un numero valido")
	}
	fmt.Printf("%d",index)
	operations.UpdateWork(index,&ListWorks)
	operations.Resume(&ListWorks)





}