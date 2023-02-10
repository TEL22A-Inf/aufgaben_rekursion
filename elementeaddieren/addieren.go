package elementeaddieren

// Liefert die Summe aller Elemente in list.
func AddElements(list []int) int {

	// Beobachtung: Die Summe aller Elemente in list ist:

	// * 0, falls die Liste leer ist
	if len(list) == 0 {
		return 0
	}

	// * Das erste Element
	// * plus die Summe aller restlichen Elemente
	head, tail := list[0], list[1:]
	//head := list[0]
	//tail := list[1:]

	return head + AddElements(tail)
}
