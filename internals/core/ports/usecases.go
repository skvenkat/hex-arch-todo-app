package ports

import (
	"github.com/skvenkat/hex-arch-todo-app/internal/core/domain"
)

type ToDoUseCase interface {
	Get(id string) (*domain.ToDo, error)
	List() ([]domain.ToDo, error)
	Create(title, description string) (*domain.ToDo, error)
}
