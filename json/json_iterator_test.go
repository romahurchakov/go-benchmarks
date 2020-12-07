package json_bench

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/romahurchakov/go-benchmarks/json/data"
)

func benchmarkUnmarshalJsoniter(data []byte, b *testing.B) {
	b.ReportAllocs()

	var err error
	for n := 0; n < b.N; n++ {
		m := make(map[string]interface{})
		err = jsoniter.Unmarshal(data, &m)
	}

	publicErr = err
}

func benchmarkMarshalJsoniter(data map[string]interface{}, b *testing.B) {
	b.ReportAllocs()

	var err error
	for n := 0; n < b.N; n++ {
		_, err = jsoniter.Marshal(data)
	}

	publicErr = err
}

func BenchmarkUnmarshalJsoniterSmall(b *testing.B) {
	benchmarkUnmarshalJsoniter(data.SmallJSON, b)
}

func BenchmarkUnmarshalJsoniterMedium(b *testing.B) {
	benchmarkUnmarshalJsoniter(data.MediumJSON, b)
}

func BenchmarkUnmarshalJsoniterLarge(b *testing.B) {
	benchmarkUnmarshalJsoniter(data.LargeJSON, b)
}

func BenchmarkUnmarshalJsoniterXLarge(b *testing.B) {
	benchmarkUnmarshalJsoniter(data.XLargeJSON, b)
}

func BenchmarkMarshalJsoniterSmall(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.SmallJSON), &m)

	b.ResetTimer()
	benchmarkMarshalJsoniter(m, b)
}

func BenchmarkMarshalJsoniterMedium(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.MediumJSON), &m)

	b.ResetTimer()
	benchmarkMarshalJsoniter(m, b)
}

func BenchmarkMarshalJsoniterLarge(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.LargeJSON), &m)

	b.ResetTimer()
	benchmarkMarshalJsoniter(m, b)
}

func BenchmarkMarshalJsoniterXLarge(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.XLargeJSON), &m)

	b.ResetTimer()
	benchmarkMarshalJsoniter(m, b)
}
