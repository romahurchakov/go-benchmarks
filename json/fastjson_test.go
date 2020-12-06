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

func benchmarkMarshalFastJSON(value *fastjson.Value, b *testing.B) {
	b.ReportAllocs()

	var data []byte
	for n := 0; n < b.N; n++ {
		data = value.MarshalTo(data)
	}

	publicData = data
}

func BenchmarkUnmarshalFastJSONSmall(b *testing.B) {
	benchmarkUnmarshalFastJSON(data.SmallJSON, b)
}

func BenchmarkUnmarshalFastJSONMedium(b *testing.B) {
	benchmarkUnmarshalFastJSON(data.MediumJSON, b)
}

func BenchmarkUnmarshalFastJSONLarge(b *testing.B) {
	benchmarkUnmarshalFastJSON(data.LargeJSON, b)
}

func BenchmarkMarshalFastJSONSmall(b *testing.B) {
	val, _ := fastjson.Parse(data.SmallJSON)
	benchmarkMarshalFastJSON(val, b)
}

func BenchmarMarshalFastJSONMedium(b *testing.B) {
	val, _ := fastjson.Parse(data.MediumJSON)
	benchmarkMarshalFastJSON(val, b)
}

func BenchmarkMarshalFastJSONLarge(b *testing.B) {
	val, _ := fastjson.Parse(data.LargeJSON)
	benchmarkMarshalFastJSON(val, b)
}
