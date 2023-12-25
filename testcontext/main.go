package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	ctx := context.WithValue(context.Background(), Foo{}, "1")
	Bar(ctx, `{"Field":1}`)
	fmt.Println(ctx.Value(Foo{}))
}

type Foo struct {
	Field int
}

func Bar(ctx context.Context, str string) {
	b := strings.NewReader(str)
	result := Foo{}
	err := json.NewDecoder(b).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx = context.WithValue(ctx, Foo{}, result)
}
