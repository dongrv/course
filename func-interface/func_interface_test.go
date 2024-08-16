package func_interface

import "testing"

func TestCallGetter(t *testing.T) {
	fn := func(key string) ([]byte, error) {
		return []byte(key), nil
	}
	CallGetter(GetterFunc(fn), "key")

	CallGetter(GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	}), "key")

	CallGetter(new(GetterStruct), "key")

}
