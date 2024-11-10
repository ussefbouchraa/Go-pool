package piscine

func Atoi(s string) int {

res:=0
sign:= 1

runes:= []rune(s)

for i:=0; i < len(runes); i++{

	if(runes[i] == '-' || runes[i] =='+'){
		if(runes[i] == '-'){
				sign *= -1
		}
		i++
	}
	if(runes[i] < '0' || runes[i] > '9'){
		return 0
	}
	for runes[i] == '0'{
		i++;
	}
res = (res *10) + int(runes[i] - '0')
}
return res * sign

}
