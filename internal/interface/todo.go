package interface

type TodoRepositoryIF interface {
	Create(todo *Todo) error
	GetByID(id *int) (*Todo, error)
	Update(todo *Todo) error
	Delete(id int) error
}
