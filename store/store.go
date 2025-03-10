package store

import (
	"errors"
	"github.com/cmj7271/go-todo-app/entity"
)

var (
	Tasks       = &TaskStore{Tasks: map[entity.TaskID]*entity.Task{}}
	ErrNotFound = errors.New("task not found")
)

type TaskStore struct {
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}
func (ts *TaskStore) Get(id entity.TaskID) (*entity.Task, error) {
	if t, ok := ts.Tasks[id]; ok {
		return t, nil
	}
	return nil, ErrNotFound
}
func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i] = t
	}
	return tasks
}
