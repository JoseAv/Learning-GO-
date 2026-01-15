package main

import (
	"basic/filter"
	"fmt"
)

func main() {
	numsInt := []int{1, 1, 3,3, 4, 5, 2, 2,10,12,15}
	result:=filter.Recibe(numsInt,filter.Filter)
	fmt.Println(result)
}