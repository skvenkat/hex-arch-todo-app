package todo

import (
	restful "github.com/emicklei/go-restful/v3"
	"github.com/skvenkat/hex-arch-todo-app/internal/core/ports"
)

type ToDoHandler struct {
	toDoUseCase ports.ToDoUseCase
}

func NewToDoHandler(todoUseCase ports.ToDoUseCase, ws *restful.WebService) *ToDoHandler {
	handler := &ToDoHandler{
		toDoUseCase: todoUseCase,
	}

	ws.Route(ws.GET("/todo/{id}").To(handler.Get).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.GET("/todo").To(handler.List).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.POST("/todo").To(handler.Create).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))

	return handler
}

func (tdh *ToDoHandler) Get(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")

	result, err := tdh.toDoUseCase.Get(id)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var todo *ToDo = &ToDo{}

	todo.FromDomain(result)
	result.WriteAsJson(todo)
}

func (tdh *ToDoHandler) List(req *restful.Request, resp *restful.Response) {
	result, err := tdh.toDoUseCase.List()
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var todos ToDoList = ToDoList{}

	todos = todos.FromDomain(result)
	resp.WriteAsJson(todos)
}

func (tdh *ToDoHandler) Create(req *restful.Request, resp *restful.Response) {
	var data = new(ToDo)
	if err := req.ReadEntity(data); err != nil {
		resp.WriteError(500, err)
		return
	}

	result, err := tdh.toDoUseCase.Create(data.Title, data.description)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var todo ToDo = ToDo{}
	todo.FromDomain(result)
	resp.WriteAsJson(todo)
}