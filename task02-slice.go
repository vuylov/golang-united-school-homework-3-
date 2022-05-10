package homework

func reverse(input []int64) (result []int64) {
	cp := make([]int64, len(input))
	result = make([]int64, len(input))
	copy(cp, input)

	if len(cp) <= 1 {
		return cp
	}

	p := len(cp) - 1
	for i := 0; i < len(cp); i++ {
		result[i] = cp[p]
		p--
	}

	return result
}
