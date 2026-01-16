package main

import (
	"basic/operation"
	"fmt"
)

func main() {
	newSlice := []int{1,2,3}
	otherSlice := []int{3,4,5,6,12}
	
	fmt.Println(newSlice)
	operation.Addelement(&newSlice, 10)
	fmt.Println(newSlice)
	operation.Addelement(&newSlice, 11)
	fmt.Println(newSlice)

	operation.Addelement(&newSlice, 12)
	fmt.Println(newSlice)


	operation.DeletElement(&newSlice,10)
	fmt.Println(newSlice)

	operation.DeletElement(&newSlice,2)
	fmt.Println(newSlice)

	fmt.Println(operation.InterseccionSlice(newSlice, otherSlice))




}