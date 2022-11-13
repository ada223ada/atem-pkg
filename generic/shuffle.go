package generic

import (
	"math/rand"
	"time"
)

func Shuffle[T any](a []T,newSeed ...bool) {
	if(nil !=newSeed && newSeed[0]){
		rand.Seed(time.Now().UnixNano())
	}
	n := len(a)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}