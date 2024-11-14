package piscine

func BasicJoin(elems []string) string {

	if elems == nil {
		return ""
	}
	var res string
	for i := 0; i < len(elems); i++ {
		res += elems[i]
	}
	return res
}
