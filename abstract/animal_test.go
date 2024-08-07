package abstract

import (
	"strings"
	"testing"
)

func TestAnimalDo(t *testing.T) {
	tigger := &Trigger{}
	yc := &YellowCattle{}
	ab := &AfricanBison{}

	AnimalDo(tigger)
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
