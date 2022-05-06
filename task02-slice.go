package homework

func reverse(input []int64) (result []int64) {

	lenSlice := len(input)
	resultNew := make([]int64, lenSlice)

	for i := lenSlice - 1; i >= 0; i-- {
		opp := lenSlice - 1 - i
		resultNew[i], resultNew[opp] = input[opp], input[i]

	}
	return resultNew
}
