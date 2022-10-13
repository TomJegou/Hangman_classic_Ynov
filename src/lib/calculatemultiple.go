package lib

/*
Function used to calculate the number of lines between each char
in order to used different type of policies with different format
We just count the number of line the first char of the file which
is the char space.
*/

func CalculateMultiple(pathtofile string) int {
	count := 1
	t := GetFileInLine(pathtofile)
	for i := 1; i < len(t); i++ {
		count2 := 0                      //count the number of spaces in a line, if the number is greater than 0 it's ok
		for j := 0; j < len(t[i]); j++ { //looping into the line to chech every char in it
			if string(t[i][j]) == " " {
				count2++
			}
			if string(t[i][j]) != " " { // if there is something else than a space, it means that we are in the other char
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
