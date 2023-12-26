package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

func main() {
	ctx := context.Background() //context.WithValue(context.Background(), Foo{}, "1")
	fmt.Println("指针1：", &ctx)
	Bar(&ctx, `{"Field":1}`)
	fmt.Println(ctx.Value(Foo{}))
	//TestContext()
}

type Foo struct {
	Field int
}

// Bar 传入指针ctx，如果传值则函数内部修改无法影响上层数据
// context.Context 虽然是接口，但是 context.Background() 返回的是对象，所以必须传入指针才能对其修改有效
func Bar(ctx *context.Context, str string) {
	b := strings.NewReader(str)
	result := Foo{}
	err := json.NewDecoder(b).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("指针2：", ctx)
	*ctx = context.WithValue(*ctx, Foo{}, result)
	fmt.Println("指针3：", ctx)
}

// TestContext 测试上线文常见用法
func TestContext() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second)) // 截止到指定时间后自动取消
	var count int
	for {
		select {
		case <-ctx.Done():
			t, _ := ctx.Deadline() // 截止时间，是否设置
			fmt.Println("time over:", t.Unix())
			ctx = context.WithValue(ctx, "time@", t.Unix()) // 给上下文设置值，用于传参
			goto end
		case <-time.Tick(time.Second):
			count++
			fmt.Println("time:", count)
		}
	}

end:
	cancel()                                  // 主动调用cancelFun析构资源
	fmt.Println("context.Err:", ctx.Err())    // 取消context产生的错误或原因，如果上下文没有被取消则返回nil
	fmt.Println("time@:", ctx.Value("time@")) // 读取上下文设置的值

	// 拓展
	backgroundCtx := context.Background()                                           //  函数返回一个非 nil 的空 Context，它没有携带任何的值，也没有取消和超时信号。通常作为根 Context 使用。
	todoCtx := context.TODO()                                                       // 函数返回一个非 nil 的空 Context，它没有携带任何的值，也没有取消和超时信号。虽然它的返回结果和 context.Background() 函数一样，但是它们的使用场景是不一样的，如果不确定使用哪个上下文时，可以使用 context.TODO()
	withValueCtx := context.WithValue(backgroundCtx, "user", "tony")                // 函数接收一个父 Context 和一个键值对 key、val，返回一个新的子 Context，并在其中添加一个 key-value 数据对
	cancelCtx, cancel := context.WithCancel(backgroundCtx)                          // 函数接收一个父 Context，返回一个新的子 Context 和一个取消函数，当取消函数被调用时，子 Context 会被取消，同时会向子 Context 关联的 Done() 通道发送取消信号，届时其衍生的子孙 Context 都会被取消。这个函数适用于手动取消操作的场景
	cancelCauseCtx, cancelErr := context.WithCancelCause(backgroundCtx)             // 是 Go 1.20 版本才新增的，其功能类似于 context.WithCancel()，但是它可以设置额外的取消原因，也就是 error 信息，返回的 cancel 函数被调用时，需传入一个 error 参数
	timeoutCtx, cancelTimeout := context.WithTimeout(backgroundCtx, 10*time.Second) // 函数的功能是一样的，其底层会调用 WithDeadline() 函数，只不过其第二个参数接收的是一个超时时间，而不是截止时间。这个函数适用于需要在一段时间后取消操作的场景

	cancel()
	cancelErr(errors.New("手动取消"))
	cancelTimeout()

	_ = context.Cause(cancelCauseCtx) // 取消原因

	_, _, _, _ = todoCtx, withValueCtx, cancelCtx, timeoutCtx

}
