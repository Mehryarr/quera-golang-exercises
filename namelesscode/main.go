package main

func AddElement(numbers *[]int, element int) {
	*numbers = append(*numbers, element)
}

func FindMin(numbers *[]int) int {

	if *numbers == nil || len(*numbers) == 0 {
		return 0
	}

	min := (*numbers)[0]
	for _, v := range *numbers {
		if v < min {
			min = v
		}
	}
	return min

}

func ReverseSlice(numbers *[]int) {
	ReverseNums := make([]int, 0)

	for i := len(*numbers) - 1; i >= 0; i-- {
		ReverseNums = append(ReverseNums, (*numbers)[i])
	}
	copy(*numbers, ReverseNums)
}

func SwapElements(numbers *[]int, i, j int) {
	if numbers == nil || len(*numbers) == 0 || i < 0 || j < 0 || i >= len(*numbers) || j >= len(*numbers) {
		return
	}
	(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
}
