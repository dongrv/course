package generictype

// Swap 函数
func Swap[T any](a, b T) (T, T) {
	return b, a
}

// Pair 结构体
type Pair[T any] struct {
	First  T
	Second T
}

func NewPair[T any](a, b T) Pair[T] {
	return Pair[T]{First: a, Second: b}
}

// Container 接口
type Container[E any] interface {
	Add(E)
	Remove() E
}

type Comparable interface {
	~int | ~float64 | ~string | ~uintptr
}

type Stack[E Comparable] struct {
	items []E
}

func NewStack[E Comparable]() *Stack[E] {
	return &Stack[E]{items: []E{}}
}

// Push 入栈
func (s *Stack[E]) Push(item E) {
	s.items = append(s.items, item)
}

// Pop 弹出栈顶
func (s *Stack[E]) Pop() (E, bool) {
	if len(s.items) == 0 {
		var zero E
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek 只读栈顶
func (s *Stack[E]) Peek() (E, bool) {
	if len(s.items) == 0 {
		var zero E
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[E]) Len() int           { return len(s.items) }
func (s *Stack[E]) Less(i, j int) bool { return s.items[i] < s.items[j] }
func (s *Stack[E]) Swap(i, j int)      { s.items[i], s.items[j] = s.items[j], s.items[i] }

/*func SortStack[E sort.Interface](s *Stack[E]) {

}*/
