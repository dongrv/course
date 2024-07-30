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

// 例1
func TestBar_Get(t *testing.T) {
	bar := new(Bar)
	fmt.Printf("bar's addr %p\n", &bar)
	fmt.Printf("bar pointer's addr %p\n", bar)
	fmt.Printf("bar pointer's value %+v\n", *bar)
}

// 例2
func TestAnalyze(t *testing.T) {
	Analyze()
}

// 例3
func TestAnalyzeSlice(t *testing.T) {
	AnalyzeSlice()
}
