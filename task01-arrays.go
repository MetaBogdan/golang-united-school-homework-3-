package homework

func average(input [15]float32) (result float32) {

	var I, Summ float32
	for v := range input {

		Summ = Summ + input[v]
		I = I + 1
	}

	return Summ / I
}
