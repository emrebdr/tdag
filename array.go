package tdag

func RemoveFromArray(element string, array []string) []string {
	var index int
	for i, e := range array {
		if e == element {
			index = i
			break
		}
	}

	return append(array[:index], array[index+1:]...)
}
