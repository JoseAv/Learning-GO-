package filter

func Recibe(intNumber []int, Filter func([]int) []int) []int {
	// Error: to many return: que necesita que le especifiques que retorno

	return Filter(intNumber)
}

func Filter(intNumber []int) []int {
	FilterWithOutSameNumbers := []int{}
	helpMap := map[int]bool{}
	for _, val := range intNumber {
		if !helpMap[val] {
			helpMap[val] = true
			FilterWithOutSameNumbers = append(FilterWithOutSameNumbers, val)
		}
	}
	return FilterWithOutSameNumbers
}