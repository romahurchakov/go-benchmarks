package json_bench

import (
	"encoding/json"
	"testing"

	"github.com/romahurchakov/go-benchmarks/json/data"
)

var (
	publicErr  error
	publicData []byte
)

func benchmarkUnmarshalStd(data string, b *testing.B) {
	b.ReportAllocs()

	var err error
	for n := 0; n < b.N; n++ {
		m := make(map[string]interface{})
		err = json.Unmarshal([]byte(data), &m)
	}

	publicErr = err
}

func benchmarkMarshalStd(data map[string]interface{}, b *testing.B) {
	b.ReportAllocs()

	var (
		err error
	)
	for n := 0; n < b.N; n++ {
		_, err = json.Marshal(data)
	}

	publicErr = err
}

func BenchmarkUnmarshalStdSmall(b *testing.B) {
	benchmarkUnmarshalStd(data.SmallJSON, b)
}

func BenchmarkUnmarshalStdMedium(b *testing.B) {
	benchmarkUnmarshalStd(data.MediumJSON, b)
}

func BenchmarkUnmarshalStdLarge(b *testing.B) {
	benchmarkUnmarshalStd(data.LargeJSON, b)
}

func BenchmarkMarshalStdSmall(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.SmallJSON), &m)

	b.ResetTimer()
	benchmarkMarshalStd(m, b)
}

func BenchmarkMarshalStdMedium(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.MediumJSON), &m)

	b.ResetTimer()
	benchmarkMarshalStd(m, b)
}

func BenchmarkMarshalStdLarge(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.LargeJSON), &m)

	b.ResetTimer()
	benchmarkMarshalStd(m, b)
}
