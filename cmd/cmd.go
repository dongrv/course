package cmd

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os/exec"
	"runtime"
)

// Example 执行命令
func Example() {
	// 使用你输入的参数，返回Cmd指针，可用于执行Cmd的方法。
	// 这里name就是我们的命令/可执行文件，例如如果要执行cmd命令，这个name就是"cmd"；
	// 如果要执行bash命令，那么这个name就是"/bin/bash"，而后面的参数args可以一个一个输入。

	// windows
	// /C 表示执行完程序后关闭命令窗口
	exec.Command(`cmd`, `/C`)
	// linux
	// -c 表示后边执行的是shell或系统命令，不加 -c 表示后续调用的是一个可执行文件或脚本
	// https://stackoverflow.com/questions/3985193/what-is-bin-sh-c
	exec.Command(`/bin/bash`, `-c`)

	cmd := exec.Command(`cmd`, `/C dir`) // 获取指令封装的结构体指针
	cmd.Dir = `c:/workspace`             // 命令在目录：c:/workspace 下执行
	/*
		这几个函数都是执行，只是在输出方式上有所区别，具体可以点击链接查看文档
		(c *Cmd) CombinedOutput() ([]byte, error) 将标准输出，错误合并输出一起返回；
		(c *Cmd) Output() ([]byte, error) 输出标准输出，错误从error返回；
		(c *Cmd) Run() error 执行任务，等待执行完成；
		(c *Cmd) Start() error， (c *Cmd) Wait() error 前者执行任务，不等待完成，用后者等待，并释放资源

		另外还提供一个标准输入/输出/错误输出的管道，我们可用这些管道中去输入输出。
		(c *Cmd) StderrPipe() (io.ReadCloser, error)
		(c *Cmd) StdinPipe() (io.WriteCloser, error)
		(c *Cmd) StdoutPipe() (io.ReadCloser, error)
	*/
	out, _ /*error*/ := cmd.CombinedOutput()                            // 返回执行的结果和报错
	gb, _ := simplifiedchinese.GB18030.NewDecoder().String(string(out)) // 解决windows环境下简体中文乱码问题
	println(gb)

	// 执行 winscp.exe 示例
	exec.Command("cmd", "/C",
		"winscp.exe",
		"/console",
		"/command",
		"option batch continue",
		"option confirm off",
		"open sftp://username:password@ip:port",
		"option transfer binary",
		"put C:\\Users\\Administrator\\Desktop\\文件名 /www/wwwroot/文件名", "exit", // 需要上传的本地文件地址，C:\xx 本地问阿金地址 /www/xx 上传到服务器的文件地址
		"/log=log_file.txt",
	)

}

// IsWinOs 是否为windows系统
func IsWinOs() bool {
	return runtime.GOOS == "windows"
}

// Command 在指定路径下执行命令
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
