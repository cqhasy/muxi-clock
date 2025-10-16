package pkg

import (
	"fmt"
	"go_clock/model/common"
	"go_clock/model/task"
	"go_clock/store"
	"go_clock/store/temp"
	"strconv"
	"time"
)

type TaskStore interface {
	CreateTask(task task.Task) error
	GetTask(id string) (task.Task, error)
	UpdateTask(id string, ta task.Task) (task.Task, error)
	DeleteTask(id string) error
	GetTaskList() ([]common.Common, error)
	GetTaskByName(name string) ([]task.Task, error)
	GetDeadLineTasks(timeStamp int64) ([]task.Task, error)
}

type TaskStoreImpl struct {
	Dao store.Store
}

func NewTaskImpl(s store.Store) *TaskStoreImpl {
	return &TaskStoreImpl{
		Dao: s,
	}
}

func (t *TaskStoreImpl) CreateTask(task task.Task) error {
	err := t.Dao.Create(store.Entity{
		Key:  strconv.FormatInt(time.Now().Unix(), 10),
		Data: task,
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskStoreImpl) GetTask(id string) (task.Task, error) {
	var ta task.Task
	re, err := t.Dao.Get(id)
	if err != nil {
		return task.Task{}, err
	}
	ta = re.Data.(task.Task)
	return ta, nil
}

func (t *TaskStoreImpl) UpdateTask(id string, ta task.Task) (task.Task, error) {
	re, err := t.Dao.Get(id)
	if err != nil {
		return task.Task{}, err
	}
	re.Data = ta
	_, err = t.Dao.Update(id, re)
	if err != nil {
		return task.Task{}, err
	}
	return task.Task{}, nil
}

func (t *TaskStoreImpl) DeleteTask(id string) error {
	err := t.Dao.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskStoreImpl) GetTaskList() ([]common.Common, error) {
	var results []common.Common
	re, err := t.Dao.Custom(func() (interface{}, error) {
		if t.Dao.Type() != store.MapStore {
			return nil, nil
		}
		dao := t.Dao.(*temp.MapStore)
		d := dao.GetTaskMapStore()
		r, err := d.List()
		if err != nil {
			return nil, err
		}
		return r, nil
	})
	if err != nil {
		return nil, err
	}
	for _, v := range re.([]task.Task) {
		results = append(results, common.Common{
			Key:  v.ID,
			Name: v.TaskName,
			Type: common.TaskType,
		})
	}
	return results, nil
}

func (t *TaskStoreImpl) GetTaskByName(name string) ([]task.Task, error) {
	if t.Dao.Type() != store.MapStore {
		return nil, fmt.Errorf("not mapstore")
	}
	d := t.Dao.(*temp.MapStore).GetTaskMapStore()
	re, err := d.GetTaskByName(name)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (t *TaskStoreImpl) GetDeadLineTasks(timeStamp int64) ([]task.Task, error) {
	if t.Dao.Type() != store.MapStore {
		return nil, fmt.Errorf("not mapstore")
	}
	d := t.Dao.(*temp.MapStore).GetTaskMapStore()
	data, err := d.List()
	if err != nil {
		return nil, err
	}
	endTime := timeStamp + 60

	var results []task.Task
	for _, v := range data {
		if v.TimeStamp > timeStamp && v.TimeStamp < endTime && v.Status == task.Planning {
			results = append(results, v)
		}
	}
	return results, nil
}
