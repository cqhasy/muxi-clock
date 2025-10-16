package temp

import (
	"fmt"
	"go_clock/model/task"
	"sync"
)

type TaskMapStore struct {
	m  map[string]task.Task
	mu sync.Mutex
}

func NewTaskMapStore() *TaskMapStore {
	return &TaskMapStore{m: make(map[string]task.Task)}
}

func (ms *TaskMapStore) Create(ta task.Task, id string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ta.ID = id
	ms.m[id] = ta
	return nil
}

func (ms *TaskMapStore) Get(id string) (task.Task, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	da, ok := ms.m[id]
	if !ok {
		return task.Task{}, fmt.Errorf("the id of %v task is null", id)
	}
	return da, nil
}

func (ms *TaskMapStore) Update(ta task.Task) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if da, ok := ms.m[ta.ID]; ok {
		if ta.AlertContent != "" {
			da.AlertContent = ta.AlertContent
		}
		if ta.TaskContent != "" {
			da.TaskContent = ta.TaskContent
		}
		if ta.TimeStamp != 0 {
			da.TimeStamp = ta.TimeStamp
		}
		da.Status = ta.Status
		ms.m[da.ID] = da
		return nil
	}
	return fmt.Errorf("the id of %v task is null", ta.ID)
}

func (ms *TaskMapStore) Delete(id string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	delete(ms.m, id)
	return nil
}

func (ms *TaskMapStore) List() ([]task.Task, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	var tasks []task.Task
	for _, v := range ms.m {
		tasks = append(tasks, v)
	}
	return tasks, nil
}

func (ms *TaskMapStore) GetTaskByName(name string) ([]task.Task, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	var tasks []task.Task
	for _, v := range ms.m {
		if v.TaskName == name {
			tasks = append(tasks, v)
		}
	}
	return tasks, nil
}
