package json_bench

import (
	"encoding/json"
	"testing"

	"github.com/romahurchakov/go-benchmarks/json/data"
	insaneJSON "github.com/vitkovskii/insane-json"
)

func benchmarkUnmarshalInsaneJSON(data string, b *testing.B) {
	b.ReportAllocs()

	var err error
	for n := 0; n < b.N; n++ {
		_, err = insaneJSON.DecodeString(data)
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

func BenchmarkMarshalInsaneJSONSmall(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.MediumJSON), &m)
	data, _ := insaneJSON.DecodeString(data.MediumJSON)

	b.ResetTimer()
	benchmarkMarshalInsaneJSON(data, b)
}

func BenchmarkMarshalInsaneJSONMedium(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.MediumJSON), &m)
	data, _ := insaneJSON.DecodeString(data.MediumJSON)

	b.ResetTimer()
	benchmarkMarshalInsaneJSON(data, b)
}

func BenchmarkMarshalInsaneJSONLarge(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.MediumJSON), &m)
	data, _ := insaneJSON.DecodeString(data.MediumJSON)

	b.ResetTimer()
	benchmarkMarshalInsaneJSON(data, b)
}
