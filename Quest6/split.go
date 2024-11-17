package piscine

func Split(s, sep string) []string {
	var result []string
	var word string

	if len(s) == 0 {
		return nil
	}
	for i := 0; i < len(s); i++ {
		if len(s) >= len(sep)+i && s[i:i+len(sep)] == sep {
			if word != "" {
				result = append(result, word)
				word = ""
			}
			i += len(sep) - 1
		} else {
			word += string(s[i])
		}
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}
