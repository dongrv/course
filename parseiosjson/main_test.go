package parseiosjson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	targetValue, testValue := Validate()
	assert.Equal(t, targetValue, testValue, "not equal")
}
