package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func main() {

	var newProductos = Producto{nombre: "Laptop",precio: 200,cantidad:3}

	updateStock := func (cantidad int)(string,int){
		newProductos.cantidad = cantidad
		return "nueva Cantidad", cantidad
	}

	calculateValue := func()float64{
		return float64(newProductos.cantidad * int(newProductos.precio))
	}

	disponible := func () string {
		if newProductos.cantidad > 0{
			return "Disponible"
		}

		return "No Disponible"
	}


	fmt.Printf("%+v :", newProductos)
	fmt.Println(updateStock(100))
	fmt.Printf("%+v :", newProductos)
	fmt.Println(calculateValue())
	fmt.Printf("%+v :", newProductos)
	fmt.Println(disponible())
	fmt.Println(updateStock(0))
	fmt.Println(disponible())



}


// nombre := "Juan"
// edad := 25
// precio := 99.99
// activo := true

// * Con %v - Go decide el formato automáticamente
// fmt.Printf("%v\n", nombre)   // Juan (lo muestra como string)
// fmt.Printf("%v\n", edad)     // 25 (lo muestra como número)
// fmt.Printf("%v\n", precio)   // 99.99 (lo muestra como float)
// fmt.Printf("%v\n", activo)   // true (lo muestra como boolean)

// * Con formatos específicos - TÚ decides
// fmt.Printf("%s\n", nombre)   // Juan (forzado a string)
// fmt.Printf("%d\n", edad)     // 25 (forzado a entero)
// fmt.Printf("%.2f\n", precio) // 99.99 (float con 2 decimales)
// fmt.Printf("%t\n", activo)   // true (%t es para boolean)


// + -> indica que nos muestre tambien los nombres en los structs solo funciona con structs