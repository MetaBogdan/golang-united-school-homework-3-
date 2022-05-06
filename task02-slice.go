package homework

func reverse(input []int64) (result []int64) {

	lenSlice := len(input)
	result := make([]int64, lenSlice)

	for i := lenSlice/2 - 1; i >= 0; i-- {
		opp := lenSlice - 1 - i
		result[i], result[opp] = input[opp], input[i]

	}
	return result
}
