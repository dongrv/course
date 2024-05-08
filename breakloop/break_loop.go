package breakloop

import (
	"math/rand"
	"time"
)

func BreakLoop() {
outerloop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			switch j % 5 {
			case 0:
				println("break")
				break outerloop
			}
		}
	}
	println("break2")
}

func BreakSwitch() {
	for i := 0; i < 10; i++ {
		switch i {
		case 5:
			break
		default:
			time.Sleep(time.Millisecond * 10)
			println(i)
		}
	}
	println("done")
}

func LoopInSwitch() {
	list := make([]int, 3)
	for i := range list {
		switch i {
		case 0:
			println("continue:", i)
			continue
		case 1:
			for j := 0; j < 10; j++ {
				if j == 2 {
					println("for break:", j)
					break
				}
				println("for break2:", j)
			}
		case 2:
			println("switch:", i)
		}
		println("i:", i)
	}
}

func For() {
	for x := 10; x >= 0; x-- {
		println("for:", x)
	}
}

func Iter() {
	// 构造迭代器
	type Iterator struct {
		Fishes []int
	}
	var iterator = func(iter *Iterator) int {
		num := 0
		if len(iter.Fishes) == 0 {
			return -1
		} else if len(iter.Fishes) == 1 {
			num = iter.Fishes[0]
			iter.Fishes = nil
		} else {
			num = iter.Fishes[0]
			iter.Fishes = iter.Fishes[1:]
		}
		return num
	}
	root := &Iterator{Fishes: rand.Perm(10)}
	for i := iterator(root); i >= 0; i = iterator(root) {
		println(i)
	}
}
