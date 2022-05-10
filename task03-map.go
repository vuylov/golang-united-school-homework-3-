package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	result = make([]string, len(input))
	keys := make([]int, len(input))

	for i := range input {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	for i := range keys {
		result = append(result, input[i])
	}

	return result
}
