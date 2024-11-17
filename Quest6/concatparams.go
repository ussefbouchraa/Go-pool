package piscine

func ConcatParams(args []string) string {
	if args == nil {
		return ""
	}
	var res string = ""
	for i := 0; i < len(args); i++ {
		res += args[i]
		if i != len(args)-1 {
			res += "\n"
		}
	}
	return res
}
