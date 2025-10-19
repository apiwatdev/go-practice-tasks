package memory

import (
	"sync"

	dom "github.com/apiwatdev/go-practice-tasks/tasks/01-crud-api/internal/domain/todo"
)

type TodoRepo struct {
	mu    sync.Mutex
	seq   int
	items map[int64]dom.Todo
}

func NewTodoRepo() *TodoRepo {
	return &TodoRepo{items: make(map[int64]dom.Todo)}
}

func (r *TodoRepo) Create(t dom.Todo) (dom.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.seq++
	t.ID = int64(r.seq)
	r.items[t.ID] = t
	return t, nil
}

func (r *TodoRepo) Get(id int64) (dom.Todo, bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.items[id]
	return t, ok, nil
}

func (r *TodoRepo) List(offset, limit int) ([]dom.Todo, int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	all := make([]dom.Todo, len(r.items))
	for _, v := range r.items {
		all = append(all, v)
	}
	total := len(all)
	if offset > total {
		return []dom.Todo{}, total, nil
	}
	end := offset + limit
	if end > total {
		end = total
	}
	return all[offset:end], total, nil
}

func (r *TodoRepo) Update(id int64, t dom.Todo) (dom.Todo, bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	old, ok := r.items[id]
	if !ok {
		return dom.Todo{}, false, nil
	}
	old.Title = t.Title
	old.IsDone = t.IsDone
	r.items[id] = old
	return old, true, nil
}

func (r *TodoRepo) Delete(id int64) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return false, nil
	}
	delete(r.items, id)
	return true, nil
}
