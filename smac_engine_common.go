package smac

func minMax(runes []rune) (rune, rune) {

	if len(runes) == 1 {
		return runes[0], runes[0]
	}

	var min, max rune

	if runes[0] > runes[1] {
		max = runes[0]
		min = runes[1]
	} else {
		max = runes[1]
		min = runes[0]
	}

	for i := 2; i < len(runes); i++ {
		if runes[i] > max {
			max = runes[i]
		} else if runes[i] < min {
			min = runes[i]
		}
	}

	return min, max
}
