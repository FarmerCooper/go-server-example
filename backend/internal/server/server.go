package server

import (
    "context"
	"time"

    todov1 "example.com/todo-platform/backend/api/todo/v1"
    "example.com/todo-platform/backend/internal/store"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type Server struct {
    todov1.UnimplementedTodoServiceServer

    store *store.Store
}

func New(s *store.Store) *Server {
    return &Server{
        store: s,
    }
}

func (s *Server) CreateTodo(ctx context.Context, request *todov1.CreateTodoRequest) (*todov1.Todo, error) {

    todo, err := s.store.Create(
		ctx,
		request.GetTitle(),
		request.GetCompleted(),
        request.GetPriority(),
	)

	if err != nil {
		return nil, status.Error(
			codes.Internal,
			"failed to create todo",
    	)
	}

	return toProto(todo), nil
}

func toProto(todo store.Todo) *todov1.Todo {
    return &todov1.Todo{
        Id:        todo.ID,
        Title:     todo.Title,
        Completed: todo.Completed,
        CreatedAt: todo.CreatedAt.Format(time.RFC3339),
        DueDate:   todo.CreatedAt.Format(time.RFC3339),
        Priority:  todo.Priority,
    }
}

func (s *Server) ListTodos(ctx context.Context, request *todov1.ListTodosRequest) (*todov1.ListTodosResponse, error) {

    todos, err := s.store.List(
        ctx,
        request.Completed,
        request.Priority,
    )

    if err != nil {
        return nil, status.Error(
            codes.Internal,
            "failed to list todos",
        )
    }

    response := &todov1.ListTodosResponse{
        Todos: make(
            []*todov1.Todo,
            0,
            len(todos),
        ),
    }

    for _, todo := range todos {
        response.Todos = append(
            response.Todos,
            toProto(todo),
        )
    }

    return response, nil
}

func (s *Server) GetTodo(ctx context.Context, request *todov1.GetTodoRequest) (*todov1.Todo, error) {

	todo, found, err := s.store.Get(
        ctx,
        request.GetId(),
    )

    if err != nil {
        return nil, status.Error(
            codes.Internal,
            "failed to get todo",
        )
    }

    if !found {
        return nil, status.Errorf(
            codes.NotFound,
            "todo %d not found",
            request.GetId(),
        )
    }

    return toProto(todo), nil
}
