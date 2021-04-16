package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// activityCmd represents the activity command
var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "activity info",
	Long:  `activity info`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("activity called top ", *top)
		if *top {
			printTopActivities()
			return
		}
	},
}

var top *bool

func init() {
	rootCmd.AddCommand(activityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// activityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	top = activityCmd.Flags().BoolP("top", "t", false, "watch top running activities")
}

func printTopActivities() {
	// <-getOutputContinually("adb shell dumpsys activity activities")
	// execCommand("adb", "shell", "dumpsys", "activity", "activities")
	str, err := runCommandWithErr("adb shell dumpsys activity activities")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		br := bufio.NewReader(strings.NewReader(str))
		flag := false
		for {
			l, e := br.ReadBytes('\n')
			if e == io.EOF {
				break
			}
			line := string(l)
			trimLine := strings.TrimSpace(line)
			if strings.HasPrefix(trimLine, "Running activities") {
				flag = true
			}
			if flag {
				fmt.Print(line)
			}
			if strings.HasPrefix(trimLine, "Run #0") {
				flag = false
			}

		}
	}
}

func getOutputContinually(name string, args ...string) <-chan struct{} {
	cmd := exec.Command(name, args...)
	closed := make(chan struct{})
	defer close(closed)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdoutPipe.Close()

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() { // 命令在执行的过程中, 实时地获取其输出
			data, err := simplifiedchinese.GB18030.NewDecoder().Bytes(scanner.Bytes()) // 防止乱码
			if err != nil {
				fmt.Println("transfer error with bytes:", scanner.Bytes())
				continue
			}

			fmt.Printf("%s\n", string(data))
		}
	}()

	if err := cmd.Run(); err != nil {
		// panic(err)
	}
	return closed
}

func execCommand(commandName string, args ...string) bool {
	cmd := exec.Command(commandName, args...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stdout.Close()

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			fmt.Println(err2)
			break
		}
		fmt.Print(line)
	}

	cmd.Wait()
	return true
}
