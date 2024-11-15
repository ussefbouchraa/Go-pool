package piscine

func Join(strs []string, sep string) string {
	if strs == nil || len(sep) == 0 {
		return ""
	}
	var res string
	for i := 0; i < len(strs); i++ {
		if i == len(strs)-1 {
			res += strs[i]
		} else {
			res += strs[i] + sep
		}
	}
	return res
}
