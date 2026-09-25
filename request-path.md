                                     [NUXT Framework] 
                                            │
                                            ├── POST :8080/todos
                                            ▼
                                        [KrakenD]
                                            |
                                            ├── Matches POST /todos
                                            |
                                            ├── Sends HTTP/JSON req to http://host.docker.internal:8081/internal/todos
                                            ▼
                                      [gRPC-Gateway]
                                            |
                                            ├── Generates gRPC-Gateway code that matches HTTP routes and decodes
                                            |
        into the generated Protobuf Go type ├── *todov1.CreateTodoRequest
                                            |
                     Performs the gRPC call ├── todo.v1.TodoService/CreateTodo
                                            |
                                            ▼
                                    [Go TodoService]
                                            |
   Dispatched RPC through TodoServiceServer ├── func (s *Server) CreateTodo(ctx context.Context, request *todov1.CreateTodoRequest) (*todov1.Todo, error) 
                                            |
            Code implementation calls store ├── s.Store.Create(ctx, request.GetTitle(), request.GetCompleted)
                                            |
                 The store creates db model ├── Todo{Title: "Understand the architecture", Completed: false}
                                            |
                                   Executes ├── s.db.WithContext(ctx).Create(&todo)
                                            ▼
                                         [GO ORM]
                                            │
                                            ├── translates the model operation into SQL
                                            ▼
                                          [SQL]
