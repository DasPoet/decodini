package bench

import "math/rand"

const seed = 2718281828

func makeMap(size int) map[int64]int {
	var rng = rand.New(rand.NewSource(seed))

	values := make(map[int64]int, size)
	for range size {
		values[rng.Int63()] = int(rng.Int63())
	}
	return values
}
