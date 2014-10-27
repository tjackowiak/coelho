package coelho

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shouldReturnTwitterClient(t *testing.T) {
	client := NewTwitterClient()
	assert.NotNil(t, &client)
}

func Test_shouldReturnTweetsFromUser(t *testing.T) {
	client := NewTwitterClient()
	results, err := client.Tweets("maneo")
	assert.Nil(t, err)
	assert.True(t, len(results) > 1)
}

func Test_shouldReturnErrorIfScreenNameIsEmpty(t *testing.T) {
	client := NewTwitterClient()
	_, err := client.Tweets("")
	assert.NotNil(t, err)
}
