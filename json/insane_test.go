package json_bench

import (
	"testing"

	"github.com/romahurchakov/go-benchmarks/json/data"
	insaneJSON "github.com/vitkovskii/insane-json"
)

func benchmarkUnmarshalInsaneJSON(data []byte, b *testing.B) {
	b.ReportAllocs()

	var err error
	for n := 0; n < b.N; n++ {
		_, err = insaneJSON.DecodeBytes(data)
	}

	publicErr = err
}

func benchmarkMarshalInsaneJSON(root *insaneJSON.Root, b *testing.B) {
	b.ReportAllocs()

	var (
		data []byte
	)
	for n := 0; n < b.N; n++ {
		data = root.Encode(nil)
	}

	publicData = data
}

func BenchmarkUnmarshalInsaneJSONSmall(b *testing.B) {
	benchmarkUnmarshalInsaneJSON(data.SmallJSON, b)
}

func BenchmarkUnmarshalInsaneJSONMedium(b *testing.B) {
	benchmarkUnmarshalInsaneJSON(data.MediumJSON, b)
}

func BenchmarkUnmarshalInsaneJSONLarge(b *testing.B) {
	benchmarkUnmarshalInsaneJSON(data.LargeJSON, b)
}

func BenchmarkUnmarshalInsaneJSONXLarge(b *testing.B) {
	benchmarkUnmarshalInsaneJSON(data.XLargeJSON, b)
}

func BenchmarkMarshalInsaneJSONSmall(b *testing.B) {
	data, _ := insaneJSON.DecodeBytes(data.SmallJSON)

	b.ResetTimer()
	benchmarkMarshalInsaneJSON(data, b)
}

func BenchmarkMarshalInsaneJSONMedium(b *testing.B) {
	data, _ := insaneJSON.DecodeBytes(data.MediumJSON)

	b.ResetTimer()
	benchmarkMarshalInsaneJSON(data, b)
}

func BenchmarkMarshalInsaneJSONLarge(b *testing.B) {
	data, _ := insaneJSON.DecodeBytes(data.LargeJSON)

	b.ResetTimer()
	benchmarkMarshalInsaneJSON(data, b)
}

func BenchmarkMarshalInsaneJSONXLarge(b *testing.B) {
	data, _ := insaneJSON.DecodeBytes(data.XLargeJSON)

	b.ResetTimer()
	benchmarkMarshalInsaneJSON(data, b)
}
