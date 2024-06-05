package testtype

import (
	"fmt"
	"unicode"
)

type T interface {
	UserId()
}

type Default struct{}

func (d *Default) UserId() int32 {
	return 100
}

type User struct {
	*Default
	Id int32
}

func (_ *User) UserId() int32 {
	return 10001
}

func Run() {
	user := &User{}
	println(user.UserId())
}

func Rune() {
	str := "你好, 世界!"
	for _, r := range str {
		fmt.Printf("%c %T\n", r, r)
		if unicode.Is(unicode.Han, r) {
			fmt.Printf("%c 是汉字\n", r)
		} else {
			fmt.Printf("%c 不是汉字\n", r)
		}
	}
	// 在这个示例中，range关键字返回的是rune类型的值，即使原始字符串是以UTF-8编码的。这样，我们就可以正确地处理每个字符，而不仅仅是每个字节。
	// 需要注意的是，虽然rune类型等价于int32，但在语义上它被明确地用来表示字符值，这有助于代码的可读性和避免误用。
}
