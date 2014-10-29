package coelho

import (
	"bufio"
	"io/ioutil"
	"testing"
)

func BenchmarkPaoloQuoteGenerator(b *testing.B) {
	paolo, _ := NewPaolo(prepareBenchmarkFile())
	for n := 0; n < b.N; n++ {
		paolo.RandomQuote()
	}
}

func prepareBenchmarkFile() string {
	f, _ := ioutil.TempFile("", "tmpquotes")
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(`[{"source":"book", "sentence":"nothing better than you"}]`)
	w.Flush()
	return f.Name()
}
