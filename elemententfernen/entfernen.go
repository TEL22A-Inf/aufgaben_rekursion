package elemententfernen

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
func RemoveElement(list []int, pos int) []int {
	if len(list) == 0 || pos < 0 {
		return list
	}

	head, tail := list[0], list[1:]
	if pos == 0 {
		return tail
	}

	result := []int{head}
	result = append(result, RemoveElement(tail, pos-1)...)

	return result

	// return append([]int{head}, RemoveElement(tail, pos-1)...)
}
