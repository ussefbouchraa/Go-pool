package piscine

func Tolower(s string) string {
	if len(s) == 0{
		return ""
	}
	bytes := []byte(s)

	for i := 0 ; i < len(bytes); i++{
		if bytes[i] >= 65 && bytes[i] <= 90{
			bytes[i] += 32 
		}
	}
	return string(bytes)
}