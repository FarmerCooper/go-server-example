Application starts
      │
      ▼
NewHandler()          ← called once
      │
      ├── NewStore()
      │       │
      │       ▼
      │    Store #1
      │
      └── NewServer(Store #1)
              │
              ▼
           Server
              │
              │ keeps a reference to Store #1
              ▼
        all HTTP requests

todo-platform/
│
├── frontend/
│   │
│   ├── src/
│   │   ├── api/
│   │   │   └── todos.ts
│   │   │
│   │   ├── components/
│   │   │   └── TodoList.vue
│   │   │
│   │   ├── App.vue
│   │   └── main.ts
│   │
│   ├── package.json
│   └── vite.config.ts
│
├── gateway/
│   │
│   ├── grpc/
│   │   └── todo.pb
│   │
│   └── krakend.json
│
└── backend/
    │
    ├── api/
    │   └── todo/
    │       └── v1/
    │           ├── todo.proto
    │           ├── todo.pb.go
    │           └── todo_grpc.pb.go
    │
    ├── cmd/
    │   └── server/
    │       └── main.go
    │
    ├── internal/
    │   ├── server/
    │   │   └── server.go
    │   │
    │   └── store/
    │       └── store.go
    │
    ├── go.mod
    └── go.sum