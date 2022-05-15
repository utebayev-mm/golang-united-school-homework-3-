package homework

func sortMapValues(input map[int]string) (result []string) {
	// Place your code here
	var keys []int
	for key := range input {
		keys = append(keys, key)
	}

	for i := 0; i < len(keys)-1; i++ {
		for j := 0; j < len(keys)-i-1; j++ {
			if keys[j] > keys[j+1] {
				keys[j], keys[j+1] = keys[j+1], keys[j]
			}
		}
	}
	for i := 0; i < len(keys); i++ {
		for key, value := range input {
			if key == keys[i] {
				result = append(result, value)
			}
		}
	}

	return
}
