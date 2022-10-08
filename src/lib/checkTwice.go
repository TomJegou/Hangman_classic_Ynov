package lib

/*
Here we check if the letter parameter is twice in a string table t,
if yes it returns true and the table otherwhise it append the letter in t and returns false
*/
func Checktwice(letter string, t []string) ([]string, bool) {
	if IsIn(t, letter) {
		return t, true
	}
	t = append(t, letter)
	return t, false
}
