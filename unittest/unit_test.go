package unittest

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"text/template"
)

// https://geektutu.com/post/quick-go-test.html

func TestPlus(t *testing.T) {
	list := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{4, 5, 11},
		{100, 200, 300},
	} // 表驱动测试
	for _, s := range list {
		if Plus(s.a, s.b) != s.expected {
			t.Fatalf("%d + %d != %d", s.a, s.b, s.expected)
		}
	}
}

// Subtests 子测试

func TestAll(t *testing.T) {
	t.Run("Plus", func(t *testing.T) {
		if Plus(1, 2) != 3 {
			t.Fail()
		}
	})
	t.Run("Sub", func(t *testing.T) {
		if Sub(10, 2) != 8 {
			t.Fail()
		}
	})
	t.Run("Mul", func(t *testing.T) {
		if Mul(1, 2) != 2 {
			t.Fail()
		}
	})
	t.Run("Div", func(t *testing.T) {
		if Div(1, 2) != 0.5 {
			t.Fail()
		}
	})
}

func TestSub(t *testing.T) {
	t.Parallel()
}

// 帮助函数

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	fmt.Println("I'm test2")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

// 基准测试 - 并发测试
func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
