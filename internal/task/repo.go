package task

import (
	"errors"
	"fmt"
	"sort"
)

var ErrTaskNotFound = errors.New("task not found")

type Repo struct {
	data map[string]Task
}

func NewRepo() *Repo {
	return &Repo{
		data: map[string]Task{
			"t_001": {
				ID:          "t_001",
				Title:       "Первая задача",
				Description: strPtr("Учебный пример"),
				Done:        false,
			},
			"t_002": {
				ID:          "t_002",
				Title:       "Вторая задача",
				Description: strPtr("GraphQL API"),
				Done:        true,
			},
		},
	}
}

func strPtr(s string) *string {
	return &s
}

func (r *Repo) ListAll() []Task {
	tasks := make([]Task, 0, len(r.data))
	for _, t := range r.data {
		tasks = append(tasks, t)
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})
	return tasks
}

func (r *Repo) GetByID(id string) (Task, error) {
	t, ok := r.data[id]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return t, nil
}

func (r *Repo) Create(title string, description *string) Task {
	id := fmt.Sprintf("t_%03d", len(r.data)+1)
	t := Task{
		ID:          id,
		Title:       title,
		Description: description,
		Done:        false,
	}
	r.data[id] = t
	return t
}

func (r *Repo) Update(id string, title *string, description *string, done *bool) (Task, error) {
	t, ok := r.data[id]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	if title != nil {
		t.Title = *title
	}
	if description != nil {
		t.Description = description
	}
	if done != nil {
		t.Done = *done
	}
	r.data[id] = t
	return t, nil
}

func (r *Repo) Delete(id string) error {
	if _, ok := r.data[id]; !ok {
		return ErrTaskNotFound
	}
	delete(r.data, id)
	return nil
}
