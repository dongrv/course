package main

func getNumber() int {
	return 42
}

func main() {
	var externalVar int
	externalVar = getNumber()
	println("External variable:", externalVar)
}

// 打印汇编
// go tool compile -S main.go
