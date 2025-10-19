package todo

import (
	"errors"

	dom "github.com/apiwatdev/go-practice-tasks/tasks/01-crud-api/internal/domain/todo"
)

type Service interface {
	Create(t dom.Todo) (dom.Todo, error)
	Get(id int64) (dom.Todo, bool, error)
	List(page, size int) ([]dom.Todo, int, int, error)
	Update(id int64, title string, isDone bool) (dom.Todo, error)
	Delete(id int64) (bool, error)
}

type service struct {
	repo dom.TodoRepository
}

func NewService(repo dom.TodoRepository) Service {
	return &service{repo: repo}
}

var ErrNotFound = errors.New("not found")

// Create implements Service.
func (s *service) Create(t dom.Todo) (dom.Todo, error) {
	return s.repo.Create(dom.Todo{Title: t.Title})
}

// Delete implements Service.
func (s *service) Delete(id int64) (bool, error) {
	return s.repo.Delete(id)
}

// Get implements Service.
func (s *service) Get(id int64) (dom.Todo, bool, error) {
	t, ok, err := s.repo.Get(id)
	if err != nil {
		return dom.Todo{}, false, err
	}
	if !ok {
		return dom.Todo{}, false, ErrNotFound
	}
	return t, true, nil
}

// List implements Service.
func (s *service) List(page int, size int) ([]dom.Todo, int, int, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}
	offset := (page - 1) * size
	items, total, err := s.repo.List(offset, size)
	if err != nil {
		return nil, 0, 0, err
	}
	totalPage := (total + size - 1) / size
	return items, totalPage, total, nil
}

// Update implements Service.
func (s *service) Update(id int64, title string, isDone bool) (dom.Todo, error) {
	t, ok, err := s.repo.Update(id, dom.Todo{Title: title, IsDone: isDone})
	if err != nil {
		return dom.Todo{}, err
	}
	if !ok {
		return dom.Todo{}, ErrNotFound
	}
	return t, nil
}
