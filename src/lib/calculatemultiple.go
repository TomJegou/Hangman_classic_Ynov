package lib

func CalculateMultiple(pathtofile string) int {
	count := 1
	t := GetFileInLine(pathtofile)
	for i := 1; i < len(t); i++ {
		count2 := 0
		for j := 0; j < len(t[i]); j++ {
			if string(t[i][j]) == " " {
				count2++
			}
			if string(t[i][j]) != " " {
				return count
			}
		}
		if count2 == 0 {
			return count
		}
		count++
	}
	return count
}
