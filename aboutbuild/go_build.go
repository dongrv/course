package main

var version string

func main() {
	println("app.version:", Version())
}

func Version() string {
	return version
}

// 编译时添加版本信息：
// go build -ldflags="-X 'main.version=1.2.3'" .

// 另外一种写入版本控制的方式是vcs，见regexp/README.md
