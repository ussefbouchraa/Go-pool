package piscine

func Index(s string, toFind string) int {
	item := rune(toFind[0])
	for i,val := range(s){
		if val == item{
			return i
		}
	}
	return -1
}