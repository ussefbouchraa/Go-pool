package piscine

func Atoi(s string) int {
	res := 0
	sign := 1
	var i int = 0

	runes := []rune(s)
	if (runes[i] == '-') || runes[i] == '+' {
		if runes[i] == '-' {
			sign *= -1
		}
		i++
	}
	for ; i < len(runes); i++ {

		if (runes[i] == '-') || runes[i] == '+' {
			i++
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}

		res = (res * 10) + int(runes[i]-'0')
	}
	return res * sign
}

func TrimAtoi(s string) int {
	res := ""
	if len(s) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' || s[i] == '-' || s[i] == '+' {
			res += string(s[i])
		}
	}
	if len(res) == 0 {
		return 0
	}

	return (Atoi(res))
}
