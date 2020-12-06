package json_bench

import (
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/romahurchakov/go-benchmarks/json/data"
)

type message map[string]interface{}

func (m message) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	var value interface{}
	if err := dec.Interface(&value); err != nil {
		return err
	}
	(m)[key] = value
	return nil
}
func (m message) NKeys() int {
	return 0
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
