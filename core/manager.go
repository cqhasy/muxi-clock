package core

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-shellwords"
	"go_clock/core/cmd"
	"go_clock/model/task/pkg"
	"go_clock/store/temp"
	"os"
)

func AppStart() {
	d := temp.NewMapStore()
	d.InitTables()
	t := pkg.NewTaskImpl(d)
	rd := cmd.NewCmdRoot()
	td := cmd.NewTaskCmd(t)
	td.AddCreateCmd()
	td.AddGetCmd()
	rd.AddTaskCmd(td)
	a := NewMuxiAlertImpl(t)
	rd.Execute()
	a.Execute()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ") // 提示符
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		parser := shellwords.NewParser()
		args, err := parser.Parse(line)
		if err != nil {
			panic(err)
		}
		rd.SetArgs(args) // 设置 cobra 执行参数
		rd.Execute()
	}
}
