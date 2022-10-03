package lib

import (
	"math/rand"
	"time"
)

func ChosseRandomWord(dic []string) string {
	return dic[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(dic))]
}
