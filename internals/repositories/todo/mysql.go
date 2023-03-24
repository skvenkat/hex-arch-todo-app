package todo

import (
	"database/sql"
	"fmt"

	"github.com/skvenkat/hex-arch-todo-app/internal/core/domain"
	"github.com/skvenkat/hex-arch-todo-app/internal/core/ports"
)

type todoMysql struct {
	ID			string
	Title		string
	Description	string
}

type todoListMysql []todoMysql

func (m *todoMysql) ToDomain() *domain.ToDo {
	return &domain.ToDo{
		ID:				m.ID,
		Title:			m.Title,
		Description:	m.Description,
	}
}

func (m todoListMysql) ToDomain() todoListMysql {
	todos := make([]domain.ToDo, len(m))
	for k, td := range todos {
		todo := td.ToDomain()
		todos[k] = *todo
	}

	return todos
}

type todoMysqlRepo struct {
	db *sql.DB
}

func NewTodoMysqlRepo(db *sql.DB) ports.ToDoRepository {
	return &todoMysqlRepo{
		db: db,	
	}
}

func (m *todoMysqlRepo) Get(id string) (*domain.ToDo, error) {
	var todo todoMysql = todoMysql{}
	sqlQuery := fmt.Sprintf("SELECT id, title, description FROM todo WHERE id = '%s'", id)
	result := m.db.QueryRow(sqlQuery)
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
		return nil, err
	}

	return todo.ToDomain(), nil
}

func (m *todoMysqlRepo) List() ([]domain.ToDo, error) {
	var todos todoListMysql
	sqlQuery := "SELECT id, title, description FROM todo"
	result, err := m.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	for result.Next() {
		todo := todoMysql{}

		if err := result.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos.ToDomain(), nil
}

func (m *todoMysqlRepo) Create(todo *domain.ToDo) (*domain.ToDo, error) {
	sqlQuery := "INSERT INTO todo (id, title, description) VALUES (?, ?, ?)"
	_, err := m.db.Exec(sqlQuery, todo.ID, todo.Title, todo.Description)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
