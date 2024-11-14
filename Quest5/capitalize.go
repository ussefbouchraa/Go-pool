package piscine

func Capitalize(s string) string {

	bytes := []byte(s)

	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 65 && bytes[i] <= 90 {
			bytes[i] += 32
		}
	}
	bytes[0] -= 32
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ' ' || bytes[i] == '+' && bytes[i+1] != 0 {
			if bytes[i + 1] >= 97 && bytes[i + 1] <= 122{
				bytes[i+1] -= 32
			}
		}
	}
	return string(bytes)
}
