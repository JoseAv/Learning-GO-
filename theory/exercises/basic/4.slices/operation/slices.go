package operation

func addelement(newSlice *[]int, value int) string {

	*newSlice = append(*newSlice, value)
	return "Create"
}

func deletElement(newSlice *[]int, value int) string {
	forArray := []int{}
	for _, num := range *newSlice {
		if num != value {
			forArray = append(forArray, num)
		}
	}
	*newSlice = forArray
	return "Delete complete"
}

func deletElementForm2(newSlice *[]int, value int) string {
	forArray := []int{}
	for index, num := range *newSlice {
		if num == value {
			// (*newSlice) -> primero desreferencias no puedes sacar indice sin esto antes
			// (*newSlice)[index+1:]... esto se hace porque append espera valores no slices
			// entonces sin ... seria -> [4,5,6] pero con puntes seria 4,5,6
			// es el spreed operator de js.
			*newSlice = append((*newSlice)[:index], (*newSlice)[index+1:]...)
		}
	}
	*newSlice = forArray
	return "Delete complete"
}