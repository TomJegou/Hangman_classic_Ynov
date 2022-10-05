package lib

/*
Here we check if the letter is in the slice of strings
for the lenght of the tab  if the string of tab of i equal the string it return true else false
*/
func IsIn(t []string, s string) bool {
	for i := 0; i < len(t); i++ {
		if string(t[i]) == s {
			return true
		}
	}
	return false
}
