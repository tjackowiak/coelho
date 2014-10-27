package coelho

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_shouldReturnErrorForNonExistingFile(t *testing.T) {
	_, err := NewPaolo("non-existing-file.json")
	assert.NotNil(t, err)
}

func Test_shouldReturnCreateNewPaolo(t *testing.T) {
	Paolo, err := NewPaolo(prepareTestFile(t))
	assert.Nil(t, err)
	assert.NotNil(t, Paolo)
}

func Test_shouldReturnRandomQuote(t *testing.T) {
	Paolo, _ := NewPaolo(prepareTestFile(t))
	randomQuote := Paolo.RandomQuote()
	assert.NotNil(t, randomQuote)
	assert.Equal(t, "nothing better than you", randomQuote.Sentence)
}

func Test_shouldReturnNewInstance(t *testing.T) {
	quotes, err := NewQuotes(prepareTestFile(t))

	assert.Nil(t, err)
	assert.NotNil(t, quotes)
}

func Test_shouldReturnQuote(t *testing.T) {
	quote, err := NewQuotes(prepareTestFile(t))

	assert.Nil(t, err)
	assert.NotNil(t, quote[0])
}

func Test_shouldReturnRightQuote(t *testing.T) {
	quote, _ := NewQuotes(prepareTestFile(t))
	assert.Equal(t, "nothing better than you", quote[0].Sentence)
}

func prepareTestFile(t *testing.T) string {
	f, err := ioutil.TempFile("", "tmpquotes")
	if err != nil {
		t.Fail()
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(`[{"book_title":"book", "sentence":"nothing better than you"}]`)
	w.Flush()
	return f.Name()
}
