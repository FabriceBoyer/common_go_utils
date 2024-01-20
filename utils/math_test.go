package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://unexpected-go.com/theres-no-min-function.html
func TestMinInt(t *testing.T) {
	assert.Equal(t, Min(5, 2), 2, "Min should be 2")
}

func TestMaxInt(t *testing.T) {
	assert.Equal(t, Max(5, 2), 5, "Max should be 5")
}

func TestClamp(t *testing.T) {
	assert.Equal(t, Clamp(15, 0, 10), 10, "Clamp value not correct")
	assert.Equal(t, Clamp(-15, 0, 10), 0, "Clamp value not correct")
	assert.Equal(t, Clamp(5, 0, 10), 5, "Clamp value not correct")
}
