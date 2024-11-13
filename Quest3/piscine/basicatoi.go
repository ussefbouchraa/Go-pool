package piscine

func BasicAtoi(s string) int {

	res:=0
	for _,val:= range(s){
		res = (res *10) + int(val - '0')
	}
	return res
}