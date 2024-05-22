package tdd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddOne(t *testing.T) {
	assert.Equal(t, 10, AddOne(9))
}

func TestTransform(t *testing.T) {
	assert.Equal(t, 100, Transform("100"))
	assert.NotEqual(t, 100, Transform("gawergg"))
}

func TestDistance(t *testing.T) {
	assert.Equal(t, 1.0, Distance(0, 0, 0, 1))
}
