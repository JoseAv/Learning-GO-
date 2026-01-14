package basico

import (
	"fmt"
	"strconv"
)

func ShowVar() {
	valueInt := 10
	ValueFloat := 10.20
	valueString := "Hola"
	valueBool := true

	fmt.Println(valueInt)
	fmt.Println(ValueFloat)
	fmt.Println(valueString)
	fmt.Println(valueBool)
}

func MoreVar(){
	const pi float64 = 3.1416
	const name string = "Jose"
	const edad int16 = 26

	fmt.Println(pi)
	fmt.Println(name)
	fmt.Println(edad)
}


func ConverData(){
	// numeros se convierte sin problema
	intToFloat := float32(24)
	value := ""
	// string necesita el convetirod srtconv
	stringToInt,errorConvert := strconv.Atoi(value)
	fmt.Println(intToFloat)
	fmt.Println(errorConvert)
	// si hay error lo deja el int a un valor a 0 o string a ""
	fmt.Println(stringToInt)
}

var (
	 Name = "New Data"
	 Age = 20
	)

func ChangeValue(){
	num1 := 10
	num2 := 20

	num1, num2 = num2 , num1
	fmt.Println(num1)
	fmt.Println(num2)
}

func CalculateArea(alto, ancho int)int{
	return alto * ancho
}

func Calculator(a,b int){

	 sum:=func (a,b int)int{
		return a+b
	}
	
	res:= func(a,b int)int {
		return a-b
	}
	
	mul:= func(a,b int)int{
		return a * b
	}
	
	div:= func(a,b float32)float32{
		return a/b
	}
	fmt.Println(sum(a,b))
	fmt.Println(res(a,b))
	fmt.Println(mul(a,b))
	fmt.Println(div(float32(a),float32(b)))

}