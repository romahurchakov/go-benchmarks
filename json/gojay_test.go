package json_bench

import (
	"encoding/json"
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/romahurchakov/go-benchmarks/json/data"
)

type message map[string]interface{}

func (m message) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	str := ""
	err := dec.String(&str)
	if err != nil {
		return err
	}
	m[k] = str
	return nil
}

func (m message) NKeys() int {
	return 0
}

func (m message) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range m {
		enc.AddInterfaceKey(k, v)
	}
}

func (m message) IsNil() bool {
	return m == nil
}

func benchmarkUnmarshalGojay(data []byte, b *testing.B) {
	b.ReportAllocs()

	var err error
	for n := 0; n < b.N; n++ {
		msg := make(message)
		err = gojay.UnmarshalJSONObject(data, &msg)
	}

	publicErr = err
}

func benchmarkMarshalGojay(msg message, b *testing.B) {
	b.ReportAllocs()

	var (
		err error
	)
	for n := 0; n < b.N; n++ {
		_, err = gojay.MarshalJSONObject(msg)
	}

	publicErr = err
}

func BenchmarkUnmarshalGojaySmall(b *testing.B) {
	benchmarkUnmarshalGojay(data.SmallJSON, b)
}

func BenchmarkUnmarshalGojayMedium(b *testing.B) {
	benchmarkUnmarshalGojay(data.MediumJSON, b)
}

func BenchmarkUnmarshalGojayLarge(b *testing.B) {
	benchmarkUnmarshalGojay(data.LargeJSON, b)
}

func BenchmarkUnmarshalGojayXLarge(b *testing.B) {
	benchmarkUnmarshalGojay(data.XLargeJSON, b)
}

func BenchmarkMarshalGojaySmall(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.SmallJSON), &m)

	b.ResetTimer()
	benchmarkMarshalGojay(m, b)
}

func BenchmarkMarshalGojayMedium(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.MediumJSON), &m)

	b.ResetTimer()
	benchmarkMarshalGojay(m, b)
}

func BenchmarkMarshalGojayLarge(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.LargeJSON), &m)

	b.ResetTimer()
	benchmarkMarshalGojay(m, b)
}

func BenchmarkMarshalGojayXLarge(b *testing.B) {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data.XLargeJSON), &m)

	b.ResetTimer()
	benchmarkMarshalGojay(m, b)
}
