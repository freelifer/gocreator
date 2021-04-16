package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"
)

// Result 结果
type Result struct {
	msg string
	err int
}

// NewResult 创建结果对象
func NewResult(msg string, err int) *Result {
	return &Result{msg, err}
}

// NotFound aa
var NotFound = &Result{"", 1}

// Runnable 执行命令回调
type Runnable func(string)

// RunnableReturn 执行命令回调
type RunnableReturn func(Result)

// pmListPackages 基础接口
// 获取当前手机下第三方应用列表
func pmListPackages(name string, run Runnable) {
	runCommand("adb shell pm list packages -3", func(result string) {
		rr := ""
		br := bufio.NewReader(strings.NewReader(result))
		for {
			l, e := br.ReadBytes('\n')
			if e == io.EOF {
				break
			}
			// package:
			trimLine := string(l[8:])
			if name != "" {
				if strings.Index(trimLine, name) != -1 {
					rr += trimLine
				}
			} else {
				rr += trimLine
			}
		}
		if rr == "" {
			fmt.Println("Not Found this packageNme")
			return
		}
		run(rr)
	})
}

// runCommand 回调式执行命令
func runCommand(cmd string, run Runnable) {
	result := ""
	var err error
	if runtime.GOOS == "windows" {
		result, err = runInWindowsWithErr(cmd)
	} else {
		result, err = runInLinuxWithErr(cmd)
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	run(result)
}

func runCommandWithErr(cmd string) (string, error) {
	if runtime.GOOS == "windows" {
		return runInWindowsWithErr(cmd)
	} else {
		return runInLinuxWithErr(cmd)
	}
}

func runInLinuxWithErr(cmd string) (string, error) {
	fmt.Println("Running Linux cmd: " + cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result)), err
}

func runInWindowsWithErr(cmd string) (string, error) {
	fmt.Println("Running Windows cmd: " + cmd)
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result)), err
}
