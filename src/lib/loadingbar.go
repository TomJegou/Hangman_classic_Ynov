package lib

import (
	"time"
)

func LoadingBar(t time.Duration) { //t = time in millisecond

	ldbar := "[░░░░░░░░░░░░░░░░░░░░░░░░░░░]"
	ldbarRune := []rune(ldbar)
	for i := 1; i < len(ldbarRune)-1; i++ {
		Clear()
		ldbarRune[i] = rune('█')
		PrintColor(string(ldbarRune), "Red")
		time.Sleep(t * time.Millisecond)
	}
	Clear()
	PrintColor(string(ldbarRune), "Green")
	time.Sleep(500 * time.Millisecond)
	Clear()
	PrintColor("Loading Complet\n", "Green")
	time.Sleep(500 * time.Millisecond)

}

/*
[░░░░░░░░░░░░░░░░░░░░░░░░░░░]
[███████████░░░░░░░░░░░░░░░░]
[███████████████████████████]

░ 176
█ 219
*/
