package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0{
		return -1
	}
	item := rune(toFind[0])
	for i,val := range(s){
		if val == item{
			return i
		}
	}
	return -1
}