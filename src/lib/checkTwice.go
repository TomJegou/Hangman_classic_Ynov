package lib

/*
Here we check if the letter as been already play if yes it return true otherwhise it's return false
*/
func Checktwice(letter string, t []string) ([]string, bool) {
	if IsIn(t, letter) { // check if the letter is in
		return t, true
	}
	t = append(t, letter) // append the letter that who just been played and are not already in the tab
	return t, false
}
