## TodoList API

A simple CRUD written in [Go](https://go.dev/), using [Gin](https://github.com/gin-gonic/gin) framework.

## Run

- Install dependencies
```sh
$ go get .
```
- Create `.env` file, check: [.env.example](.env.example)
- Start application
```sh
$ go run server.go
```
## Available routes

- POST - ``api/v1/todo`` - Create Todo
- GET - ``api/v1/todos`` - Todo List
- GET - ``api/v1/todo/{id}`` - Todo
- DELETE - ``api/v1/todo/{id}`` - Delete Todo
- UPDATE - ``api/v1/todo/{id}`` - Update Todo

