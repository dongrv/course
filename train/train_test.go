package train

import (
	"fmt"
	"testing"
)

func TestCausePanic(t *testing.T) {
	CausePanic()
}

func TestDiv(t *testing.T) {
	c := Div(1, 0)
	t.Log(c)
}

func TestBar_Counter(t *testing.T) {
	//deadlock test
	//bar := &Bar{}
	//t.Log(bar.Counter())
}

func TestBar_Get(t *testing.T) {
	bar := new(Bar)
	fmt.Printf("bar's addr %p\n", &bar)
	fmt.Printf("bar pointer's addr %p\n", bar)
	fmt.Printf("bar pointer's value %+v\n", *bar)
}

func TestAnalyze(t *testing.T) {
	Analyze()
}

func TestAnalyzeSlice(t *testing.T) {
	AnalyzeSlice()
}

func TestPointerAndEscape(t *testing.T) {
	PointerAndEscape()
}

func TestReflectMethod(t *testing.T) {
	ReflectMethod()
}

func TestDeleteSlice(t *testing.T) {
	t.Helper()
	DeleteSlice()
}

func TestSliceReference(t *testing.T) {
	SliceReference()
}

func TestAccessUnderlyingSlice(t *testing.T) {
	AccessUnderlyingSlice()
}

func TestParseFuncStack(t *testing.T) {
	ParseFuncStack()
}

func TestSwitchCase(t *testing.T) {
	SwitchCase()
}
