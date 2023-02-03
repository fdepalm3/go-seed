package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	AssertionString string `json:"assertion_string"`
}

func TestReadJsonFile(t *testing.T) {

	t.Run("Load test config from file", func(t *testing.T) {
		config := &TestConfig{}

		err := ReadJsonFile("json_test.json", config)

		assert.NoError(t, err)
		assert.NotEmpty(t, config.AssertionString)
	})

	t.Run("Can't load config from non-existant file", func(t *testing.T) {
		config := &TestConfig{}

		err := ReadJsonFile("non_existant_file.json", config)

		assert.Error(t, err)
	})

}
