package checksvn

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// 一些标准操作方法

type At string

func (a At) String() string {
	return string(a)
}

// StdError 标准错误
func StdError(err error, at At) {
	if err == nil {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, "%s\t%s\t错误：%s\n", time.Now().Format(time.DateTime+".000"), at.String(), err.Error())
}

// StdOut 标准输出
func StdOut(b []byte, at At) {
	_, _ = fmt.Fprintf(os.Stdout, "%s\t%s\t输出：%s\n", time.Now().Format(time.DateTime+".000"), at, b)
}
func TrimSlice(strs []string) []string {
	var result []string
	for _, str := range strs {
		result = append(result, strings.TrimSpace(str))
	}
	return result
}

// IsWinOs 是否为windows系统
func IsWinOs() bool {
	return runtime.GOOS == "windows"
}

func Command(dir string, args ...string) (string, error) {
	name := "/bin/bash"
	c := "-c"
	// 命令兼容系统差异
	if IsWinOs() {
		name = "cmd"
		c = "/C"
	}
	// 组装可执行命令结构
	args = append([]string{c}, args...)
	cmd := exec.Command(name, args...)
	if dir != `` {
		cmd.Dir = dir
	}
	// 创建获取输出命令的管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err = cmd.Start(); err != nil {
		return "", err
	}
	// 读取管道
	bytes, newErr := io.ReadAll(stdout)
	if newErr != nil {
		return "", newErr
	}
	if err = cmd.Wait(); err != nil {
		return "", err
	}
	if IsWinOs() {
		output, err := simplifiedchinese.GB18030.NewDecoder().String(string(bytes))
		if err != nil {
			return "", err
		}
		return output, nil
	}
	return string(bytes), nil
}
