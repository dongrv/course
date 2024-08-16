package func_interface

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (g GetterFunc) Get(key string) ([]byte, error) {
	return g(key)
}

type GetterStruct struct{}

func (g GetterStruct) Get(key string) ([]byte, error) {
	return nil, nil
}

func CallGetter(g Getter, key string) {
	println(g.Get(key))
}
