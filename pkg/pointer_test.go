package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPointer(t *testing.T) {
	pointer := ToPointer("asd")

	assert.NotNil(t, pointer)
	assert.Equal(t, "asd", *pointer)
}
