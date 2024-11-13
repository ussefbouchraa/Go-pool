package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if nb%2 == 0 {
		nb++
	}
	for !IsPrime(nb) {
		nb += 2
	}
	return nb
}
