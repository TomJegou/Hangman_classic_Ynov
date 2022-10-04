package lib

func Checktwice(letter string, t []string) ([]string, bool) {
	if IsIn(t, letter) {
		return t, true
	}
	t = append(t, letter)
	return t, false
}
