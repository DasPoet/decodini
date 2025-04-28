package bench

import (
	"testing"

	"github.com/lukasl-dev/decodini/pkg/decodini"
)

func BenchmarkDecodini_Transmute_MapIntoMap_100k(b *testing.B) {
	const size = 100_000

	benchmarkDecodini_Transmute[map[int64]int](b, makeMap(size))
}

func benchmarkDecodini_Transmute[Target any](b *testing.B, value any) {
	for b.Loop() {
		_, err := decodini.Transmute[Target](nil, value)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
