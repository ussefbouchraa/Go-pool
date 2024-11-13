package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 1 {
		return nb
	}

	res := 1
	for i := 0; power > i; i++ {
		res *= nb
	}

	return res
}
