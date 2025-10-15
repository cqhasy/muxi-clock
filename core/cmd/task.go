package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go_clock/model/task"
	"go_clock/model/task/pkg"
	"strconv"
	"time"
)

type TaskCmd struct {
	cmd *cobra.Command
	db  pkg.TaskStore
}

func NewTaskCmd(d pkg.TaskStore) *TaskCmd {
	return &TaskCmd{
		cmd: &cobra.Command{
			Use:   "task",
			Short: "controller task",
			Long:  "add,update,delete,get task",
		},
		db: d,
	}
}

func (tc *TaskCmd) AddCreateCmd() {
	c := &cobra.Command{
		Use:   "create",
		Short: "create task",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			content, _ := cmd.Flags().GetString("content")
			alert, _ := cmd.Flags().GetString("alert")
			timeStr, _ := cmd.Flags().GetString("time")
			if name == "" || content == "" || timeStr == "" {
				return fmt.Errorf("缺少参数，请使用 --name --content")
			}
			fmt.Println(name, content, timeStr, alert)

			if alert == "" {
				alert = "muxi小管家提醒您，您的待办事项：" + name + "需要做啦🙂"
			}
			layout := "2006-01-02 15:04"
			loc, _ := time.LoadLocation("Asia/Shanghai") // 使用北京时间
			taskTime, err := time.ParseInLocation(layout, timeStr, loc)
			if err != nil {
				return fmt.Errorf("时间格式错误，请使用格式：2025-10-14 20:00")
			}
			err = tc.db.CreateTask(task.Task{
				ID:           strconv.FormatInt(time.Now().Unix(), 10),
				TaskName:     name,
				TaskContent:  content,
				AlertContent: alert,
				TimeStamp:    taskTime.Unix(),
			})
			if err != nil {
				return fmt.Errorf("db create task err: %v", err)
			}
			return nil
		},
	}
	c.Flags().StringP("name", "n", "", "任务名称")
	c.Flags().StringP("content", "c", "", "任务内容")
	c.Flags().StringP("alert", "a", "", "提醒内容")
	c.Flags().StringP("time", "t", "", "任务时间，格式如 2025-10-14 20:00")
	tc.cmd.AddCommand(c)
}

func (tc *TaskCmd) AddGetCmd() {
	c := &cobra.Command{
		Use:   "get",
		Short: "Get task",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				return fmt.Errorf("请使用--name输入name")
			}
			tasks, err := tc.db.GetTaskByName(name)
			if err != nil {
				return fmt.Errorf("db get task err: %v", err)
			}
			for _, t := range tasks {
				fmt.Println(t)
			}
			return nil
		},
	}
	c.Flags().StringP("name", "n", "", "任务名称")
	tc.cmd.AddCommand(c)
}
