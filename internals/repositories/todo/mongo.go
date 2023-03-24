package todo

import (
	"context"

	"github.com/skvenkat/hex-arch-todo-app/internal/core/domain"
	"github.com/skvenkat/hex-arch-todo-app/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoMongo struct {
	ID			string `bson:"_id"`
	Title		string `bson:"title"`
	Description	string `bson:"desc"`
}

type todoListMongo []todoMongo

func (m *todoMongo) FromDomain(todo *domain.ToDo) {
	if m == nil {
		m = &todoMongo{}
	}

	m.ID = todo.ID
	m.Title = todo.Title
	m.Description = todo.Description
}

func (m *todoMongo) ToDomain() *domain.Todo {
	return &domain.ToDo{
		ID:				m.ID,
		Title:			m.Title,
		Description:	m.Description,
	}
}

func (m todoListMongo) ToDomain() []domain.ToDo {
	todos := make([]domain.ToDo, len(m))
	for k, td := range m {
		todo := td.ToDomain()
		todos[k] = *todo
	}

	return todos

}

type todoMongoRepo struct {
	col *mongo.Collection
}

func NewToDoMongoRepo(db *mongo.Database) ports.ToDoRepository {
	return &todoMongoRepo{
		col: db.Collection("todo"),
	}
}

func (m *todoMongoRepo) Get(id string) (*domain.ToDo, error) {
	var todo todoMongo
	result := m.col.FindOne(context.Background(), bson.M{"_id": id})

	if err := result.Decode(&todo); err != nil {
		return nil, err
	}

	return todo.ToDomain(), nil
}

func (m *todoMongoRepo) List() ([]domain.ToDo, error) {
	var todos todoListMongo
	result, err := m.col.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err := result.All(context.Background(), &todos); err != nil {
		return nil, err
	}

	return todos.ToDomain(), nil
}

func (m *todoMongoRepo) Create(todo *domain.ToDo) (*domain.ToDo, error) {
	var tdMongo *todoMongo = &todoMongo{}
	tdMongo.FromDomain(todo)

	_, err := m.col.InsertOne(context.Background(), tdMongo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
