package piscine

func ToUpper(s string) string {
	bytes := []byte(s)

	for i := 0 ; i < len(bytes); i++{
		if bytes[i] >= 97 && bytes[i] <= 122{
			bytes[i] -= 32 
		}
	}
	return string(bytes)
}