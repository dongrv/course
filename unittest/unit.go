package unittest

import (
	"errors"
	"strings"
)

func Plus(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) float64 {
	return float64(a) / float64(b)
}

const ErrServerReturn503 = "503: The service is currently unavailable" // Google验证服务不可用

func TestError() bool {
	err := errors.New("googleapi: Error 503: The service is currently unavailable., backendError")
	if err != nil && strings.Contains(err.Error(), ErrServerReturn503) {
		return true
	}
	return false
}
