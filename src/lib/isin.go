package lib

func IsIn(t []string, s string) bool {
	for i := 0; i < len(t); i++ {
		if string(t[i]) == s {
			return true
		}
	}
	return false
}
