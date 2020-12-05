package json_bench

import (
	"testing"

	"github.com/romahurchakov/go-benchmarks/json/data"
	"github.com/valyala/fastjson"
)

func benchmarkUnmarshalFastJSON(data string, b *testing.B) {
	b.ReportAllocs()

	var err error
	for n := 0; n < b.N; n++ {
		_, err = fastjson.Parse(data)
	}

	publicErr = err
}

func BenchmarkUnmarshalFastJSONSmall(b *testing.B) {
	benchmarkUnmarshalStd(data.SmallJSON, b)
}

func BenchmarkUnmarshalFastJSONMedium(b *testing.B) {
	benchmarkUnmarshalStd(data.MediumJSON, b)
}

func BenchmarkUnmarshalFastJSONLarge(b *testing.B) {
	benchmarkUnmarshalStd(data.LargeJSON, b)
}
