package bench

import (
	"maps"
	"testing"

	"github.com/mitchellh/mapstructure"
)

func BenchmarkMapstructure_EncodeDecode_MapIntoMap_100k(b *testing.B) {
	const size = 100_000

	result := make(map[int64]int)
	benchmarkMapstructure_EncodeDecode_MapIntoMap(b, makeMap(size), &result)
}

func BenchmarkGo(b *testing.B) {
	const size = 100_000

	src := makeMap(size)

	for b.Loop() {
		target := maps.Clone(src)
		if target == nil {
			panic("nil")
		}
	}
}

func benchmarkMapstructure_EncodeDecode_MapIntoMap(b *testing.B, src, dst any) {
	for b.Loop() {
		if err := mapstructure.Decode(src, dst); err != nil {
			b.Fatal(err.Error())
		}
	}
}
