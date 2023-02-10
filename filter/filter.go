package filter

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
func FilterLess(list []int, key int) []int {
	if len(list) == 0 {
		return list
	}

	head, tail := list[0], list[1:]
	result := []int{}
	// head übernehmen oder eben nicht, je nach seinem Wert
	if head <= key {
		result = append(result, head)
	}

	result = append(result, FilterLess(tail, key)...)

	return result
}

// Liefert eine Liste mit allen Elementen aus list, die größer als key sind.
func FilterGreater(list []int, key int) []int {
	// TODO
	return nil
}
