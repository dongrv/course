package manage_process

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

// CallProgram 调用程序
func CallProgram() {
	binary, err := exec.LookPath("./manage_process.exe") // LookPath 获得绝对地址；参数可以是绝对路径或相对路径。
	if err != nil {
		panic(err)
	}
	args := []string{""} // 参数输入
	env := os.Environ()  // 使用当前进程的环境变量
	err = syscall.Exec(binary, args, env)
	if err != nil {
		panic(err)
	}
}

func TestCallProgram() {
	println(`before`)
	println(syscall.Getpid()) // 当前进程ID

	// windows 不支持syscall函数
	if runtime.GOOS == "windows" {
		return
	}
	CallProgram()

	println(`after`) // 当前进程ID

}

/*
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
此方法将使用 fork & exec 的方式产生一个子进程来运行新的程序，是父子进程关系。
子进程是以守护进程的方式运行。

并且可以通过返回的 Process 对象控制子进程：
func (*Process) Kill 杀死进程
func (*Process) Release 释放进程资源
func (*Process) Signal 向进程发送信号
func (*Process) Wait 等待子进程退出，回收子进程，防止出现僵尸进程


*/
