package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveSlice(t *testing.T) {
	a := []int{1, 2, 3, 5}
	a = Remove(a, 2)
	assert.Equal(t, len(a), 3, "Array size not correct")
}
