package sns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanInput(t *testing.T) {
	str := "Hello World"
	expected := []string{"hello", "world"}
	assert.Equal(t, expected, cleanInput(str))
}
