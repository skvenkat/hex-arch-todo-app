package usecases

import (
	"github.com/skvenkat/hex-arch-todo-app/helpers"
	"github.com/skvenkat/hex-arch-todo-app/helpers/logging"
	"github.com/skvenkat/hex-arch-todo-app/internal/core/ports"
	"github.com/skvenkat/hex-arch-todo-app/internals/core/domain"
)

var (
	log = logging.NewLogger()
)

type toDoUseCase struct {
	todoRepo ports.ToDoRepository
}

func NewToDoUseCase(todoRepo ports.ToDoRepository) ports.ToDoUseCase {
	return &toDoUseCase{
		todoRepo: todoRepo,
	}
}

func (t *toDoUseCase) Get(id string) (*domain.ToDo, error) {
	todo, err := t.todoRepo.Get(id)
	if err != nil {
		log.Errorw("Error getting a ToDo from repo", logging.KeyID, id, logging.KeyError, err)
		return nil, err
	}

	return todo, nil
}

func (t *toDoUseCase) List() ([]domain.ToDo, error) {
	todoList, err := t.todoRepo.List()
	if err != nil {
		log.Errorw("Error getting ToDoList from Repo", logging.KeyError, err)
		return nil, err
	}

	return todoList, err

}

func (t *toDoUseCase) Create(title, description string) (*domain.ToDo, error) {
	todo := domain.NewToDo(helpers.RandomUuidAsString(), title, description)

	_, err := t.todoRepo.Create(todo)
	if err != nil {
		log.Errorw("Error creating a new todo from repo", "todo", logging.KeyError, err)
		return nil, err
	}

	return todo, nil
}
