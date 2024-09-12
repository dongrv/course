package main

import "regexp"

func main() {
	pattern := `(^.([\w|\.|\-]+){3,20}@[\w]{3,10}\.[a-zA-Z]{2,3}$)|(^<div[\s]+class="\w+">.*</div>)`
	reg, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	tests := []string{
		`welcom-to.10086-call@company.com`,
		`<div class="color">Hello world</div>`,
	}
	for _, v := range tests {
		if !reg.MatchString(v) {
			println("match error")
			break
		}
	}
	println("ok")
}

// go build -buildvcs=false .
// go run -buildvcs=false .
