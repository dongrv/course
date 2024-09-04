package example1

import "testing"

func TestProto(t *testing.T) {
	t.Run("inner-proto", func(t *testing.T) {
		InnerProto()
	})
	t.Run("binary-proto-marshal", func(t *testing.T) {
		MarshalBinary()
	})
	t.Run("binary-proto-unmarshal", func(t *testing.T) {
		UnmarshalBinary()
	})

}
