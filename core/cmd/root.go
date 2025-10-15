package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type RootCmd struct {
	cmd *cobra.Command
}

func NewCmdRoot() *RootCmd {
	return &RootCmd{cmd: &cobra.Command{
		Use:   "MUXI-CLOCK",
		Short: "the muxi own task clock",
		Long:  "the muxi own task clock,you can add,update,delete your own task,it will inform you when time out by cmd",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("欢迎使用muxi-clock")
		},
	}}
}

func (r *RootCmd) AddTaskCmd(t *TaskCmd) *RootCmd {
	r.cmd.AddCommand(t.cmd)
	return r
}

func (r *RootCmd) Execute() {
	if err := r.cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func (r *RootCmd) SetArgs(args []string) {
	r.cmd.SetArgs(args)
}
