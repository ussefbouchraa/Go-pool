package piscine

func StrRev(s string) string {

	var res string
	for i:=len(s) - 1 ; i >= 0; i--{
		res+= string(s[i])
	}
	return res
}