package piscine

func BasicAtoi2(s string) int {

res:=0

for _,val:=range(s){
	if val < '0' || val > '9'{
		return 0
	}
	res = (res *10) + int(val - '0')
}
return res
}

