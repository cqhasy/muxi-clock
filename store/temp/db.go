package temp

import (
	"go_clock/model/task"
	"go_clock/store"
)

type MapStore struct {
	TaskMap *TaskMapStore
}

func NewMapStore() *MapStore {
	return &MapStore{}
}

func (m *MapStore) Type() store.Type {
	return store.MapStore
}

func (m *MapStore) Connect(addr string, args ...interface{}) (interface{}, error) {
	return nil, nil
}

func (m *MapStore) InitTables(name ...string) error {
	t := NewTaskMapStore()
	m.TaskMap = t
	return nil
}

func (m *MapStore) GetTaskMapStore() *TaskMapStore {
	if m.TaskMap == nil {
		m.InitTables()
	}
	return m.TaskMap
}

func (m *MapStore) Get(key string) (store.Entity, error) {
	return store.Entity{}, nil
}

func (m *MapStore) Create(val store.Entity) error {
	err := m.GetTaskMapStore().Create(val.Data.(task.Task), val.Key)
	if err != nil {
		return err
	}
	return nil
}

func (m *MapStore) Delete(key string) error {
	return nil
}

func (m *MapStore) Update(key string, val store.Entity) (store.Entity, error) {
	return store.Entity{}, nil
}

func (m *MapStore) Custom(fn func() (interface{}, error)) (interface{}, error) {
	data, err := fn()
	return data, err
}
