package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {

	var resultNew []string

	type key_value struct {
		Key   int
		Value string
	}

	var sorted_struct []key_value

	for key, value := range input {
		sorted_struct = append(sorted_struct, key_value{key, value})
	}

	sort.Slice(sorted_struct, func(i, j int) bool {
		return sorted_struct[i].Key < sorted_struct[j].Key
	})

	for _, key_value := range sorted_struct {
		resultNew = append(resultNew, key_value.Value)
	}

	return resultNew
}
