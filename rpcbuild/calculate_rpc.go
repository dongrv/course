package rpcbuild

import (
	"errors"
	"fmt"
	"math"
	"net"
	"net/rpc"
	"syscall"
	"time"
)

// 升级版，增加接口约束

const CalculateServiceName = "path/to/HelloService" // 服务器名

// CalculateServerInterface 接口约束标准
type CalculateServerInterface = interface {
	Calculate(input Input, output *float64) error // 方法固定格式：Func(x X,y *Y) error
}

type Input struct {
	Operator Operator
	Number1  float64
	Number2  float64
}

// Expression 表达式
func (in Input) Expression() string {
	var op string
	switch in.Operator {
	case Plus:
		op = "+"
	case Sub:
		op = "-"
	case Multiple:
		op = "*"
	case Div:
		op = "/"
	}
	return fmt.Sprintf("%f %s %f", in.Number1, op, in.Number2)
}

// RegisterService 注册服务
func RegisterService(srv CalculateServerInterface) error {
	return rpc.RegisterName(CalculateServiceName, srv)
}

type CalculateService struct{}

func (c *CalculateService) Calculate(input Input, output *float64) error {
	switch input.Operator {
	case Plus:
		*output = input.Number1 + input.Number2
	case Sub:
		*output = input.Number1 - input.Number2
	case Multiple:
		*output = input.Number1 * input.Number2
	case Div:
		if input.Number2 == 0 {
			return errors.New("division by zero")
		}
		*output = input.Number1 / input.Number2
	}
	return nil
}

func CalculateRun() {
	if err := RegisterService(&CalculateService{}); err != nil {
		exit("Error registering", err)
	}
	listener, err := net.Listen("tcp", ListenAddr)
	if err != nil {
		exit("Error listening", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(*net.OpError); ok && (ne.Temporary() || ne.Timeout()) {
				continue // 临时错误或网络超时
			} else if _, ok := err.(*syscall.Errno); ok {
				continue // Linux系统调用错误
			}
			exit("Error accepting", err) // 其他错误退出
		}
		go rpc.ServeConn(conn)
	}

}

type CalculateClient struct {
	*rpc.Client
}

func (c *CalculateClient) Connect(addr string) error {
	ln, err := rpc.Dial("tcp", addr)
	if err != nil {
		return err
	}
	c.Client = ln
	return nil
}

func (c *CalculateClient) FormatMethod(method string) string {
	return CalculateServiceName + "." + method
}
func (c *CalculateClient) Calculate(input Input, output *float64) error {
	// 异步调用 c.Go(c.FormatMethod("Calculate"), input, output,<-ch)
	return c.Call(c.FormatMethod("Calculate"), input, output)
}

func CalculateDial() {
	client := &CalculateClient{}
	err := client.Connect(ListenAddr)
	if err != nil {
		exit("client connect", err)
	}
	defer client.Close()
	inputs := []Input{
		{Operator: Plus, Number1: 1, Number2: 2},
		{Operator: Sub, Number1: 10000, Number2: 1234},
		{Operator: Multiple, Number1: 9999999, Number2: 9999999},
		{Operator: Div, Number1: math.E, Number2: 0.001},
	}
	for _, input := range inputs {
		var output float64
		err = client.Calculate(input, &output)
		if err != nil {
			exit("client calculate", err)
		}
		fmt.Printf("%s = %f\n", input.Expression(), output)
		time.Sleep(time.Second)
	}
}
