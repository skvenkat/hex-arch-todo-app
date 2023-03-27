FROM golang:1.20-alpine

WORKDIR /opt/todo-app/

COPY . /opt/todo-app/
RUN go mod download

RUN go build -o todoapp migration/main.go 

EXPOSE 8080

ENTRYPOINT ["/opt/todo-app/todoapp"]
