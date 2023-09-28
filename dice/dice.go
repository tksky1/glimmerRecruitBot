package dice

import (
	"math/rand"
)

func generateRandom(mx int64) int64 {
	return rand.Int63() % mx
}

func dice(n int64, d int64) (int64, []int64) {
	res := int64(0)
	series := make([]int64, n)

	for i := int64(0); i < n; i++ {
		series[i] = generateRandom(d)

		res += series[i]
	}

	return res, series
}
