package main

import (
	"fmt"
	"shapes/operations"
)

func main() {

	circle := operations.Circle{Radio: 20}
	fmt.Println(circle.Area())
	fmt.Println(circle.Perimeter())

	triangule := operations.Triangule{Base: 2,Hight: 3}
	fmt.Println(triangule.Area())
	fmt.Println(triangule.Perimeter())

	rectangule := operations.Rectangule{Width: 2,Length: 3}
	fmt.Println(rectangule.Area())
	fmt.Println(rectangule.Perimeter())

}

