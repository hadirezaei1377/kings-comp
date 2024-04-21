package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDTypeValue(t *testing.T) {
	assert.Equal(t, "type", ID("type:val").Type())
	assert.Equal(t, "val", ID("type:val").ID())
}
