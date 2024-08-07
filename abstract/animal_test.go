package abstract

import (
	"strings"
	"testing"
)

// go test .\abstract\ -v -short -run 'TestDerived_Call|TestLongAnimalDo'
// -short 会跳过有testing.Short()方法的测试单元，一般为执行耗时或依赖硬件的测试

func TestLongAnimalDo(t *testing.T) {
	if testing.Short() {
		t.Skipf("skipping test in short mode.")
	}
	tiger := &Tiger{}
	yc := &YellowCattle{}
	ab := &AfricanBison{}

	AnimalDo(tiger)
	println(strings.Repeat("-", 10))
	AnimalDo(yc)
	println(strings.Repeat("-", 10))
	AnimalDo(ab)
}

func TestBullDo(t *testing.T) {
	yc := &YellowCattle{}
	ab := &AfricanBison{}

	BullDo(yc)
	println(strings.Repeat("-", 10))
	BullDo(ab)
}

func TestName(t *testing.T) {
	SayHandler(EatHandler(RunHandler(Name()))).Do()
}

func TestDerived_Call(t *testing.T) {
	Call()
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Call()
	}
}

// 并行运行
// go test -v -parallel=2  .\abstract\
// 重复测试
// go test -v -count=5 .\abstract\
// 生成覆盖率报告
// go test .\abstract\ -coverprofile='coverage.out'
// 查看覆盖率报告
// go tool cover -html='.\coverage.out'
// 设置测试超时时间
// go test -timeout=10s .\abstract\
// 基准测试
// go test -bench='BenchmarkCall' -v  .\abstract\
// go test -benchtime=10s
// 启用数据竞争（requires cgo; enable cgo by setting CGO_ENABLED=1）
// go test -race .\abstract\
// 同时进行单元测试和基准测试
// go test -run 'TestDerived_Call'  -bench='BenchmarkCall' -v  .\abstract\
