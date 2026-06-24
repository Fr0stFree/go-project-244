package parser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseError(t *testing.T) {
	err := newParseError(errors.New("invalid json"), "file.json")
	assert.Equal(t, `unable to parse file "file.json": invalid json`, err.Error())
}
