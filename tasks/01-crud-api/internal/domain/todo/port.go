package todo

type TodoRepository interface {
	Create(t Todo) (Todo, error)
	Get(id int64) (Todo, bool, error)
	List(offset, limit int) ([]Todo, int, error)
	Update(id int64, t Todo) (Todo, bool, error)
	Delete(id int64) (bool, error)
}
