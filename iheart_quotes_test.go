package coelho

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_retrieval(t *testing.T) {
	result, _ := NewHeartQuotes().RandomHeartQuote()
	assert.NotNil(t, result)
}

func Test_quoteCleaning(t *testing.T) {
	fakeFetcher := func() (string, error) {
		return "example &quot;content&quot;", nil
	}

	quotes := &HeartQuotes{QuoteFetcher: fakeFetcher}

	result, _ := quotes.RandomHeartQuote()
	assert.Equal(t, "example \"content\"", result)
}

func Test_shouldRemoveNewLine(t *testing.T) {
	fakeFetcher := func() (string, error) {
		return "example\n with new line", nil
	}

	quotes := &HeartQuotes{QuoteFetcher: fakeFetcher}

	result, _ := quotes.RandomHeartQuote()
	assert.Equal(t, "example with new line", result)

}
