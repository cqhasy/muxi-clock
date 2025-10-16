package core

import (
	"fmt"
	"go_clock/assets"
	"go_clock/model/task"
	"go_clock/model/task/pkg"
	"log"
	"os/exec"
	"time"
)

const (
	AlertTicker uint = 10 //second
	Retry            = 5
)

type MuxiAlertImpl struct {
	Obj pkg.TaskStore
}

func NewMuxiAlertImpl(obj pkg.TaskStore) *MuxiAlertImpl {
	return &MuxiAlertImpl{Obj: obj}
}

func (m *MuxiAlertImpl) Execute() {
	go func() {
		for {
			now := time.Now().Unix()
			ts, err := m.Obj.GetDeadLineTasks(now)
			if err != nil {
				log.Println("get deadline tasks err:", err)
			}
			for _, ta := range ts {
				go func(t task.Task) {
					for i := 0; i < Retry; i++ {
						m.alert(t.TaskName, t.AlertContent)
						time.Sleep(time.Duration(AlertTicker) * time.Second)
					}
				}(ta)
				ta.Status = task.Finished
				_, err = m.Obj.UpdateTask(ta.ID, ta)
				if err != nil {
					log.Println("update task err:", err)
				}
			}
			time.Sleep(time.Duration(AlertTicker) * time.Second)
		}
	}()
}

func (m *MuxiAlertImpl) alert(title, message string) {
	psCommand := fmt.Sprintf(`New-BurntToastNotification -Text "%s", "%s" -AppLogo "%s"`,
		title, message, assets.MuxiLogo)
	cmd := exec.Command("powershell", "-Command", psCommand)
	cmd.Run()
}
