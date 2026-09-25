                        PUBLIC SIDE
                  HTTP + JSON (REST-like)

┌────────────────────┐
│                    │
│      Vue 3 UI      │
│   localhost:5173   │
│                    │
└─────────┬──────────┘
          │
          │ POST /todos
          │ GET  /todos
          │ JSON
          ▼
┌────────────────────┐
│                    │
│      KrakenD       │
│   localhost:8080   │
│                    │
│     API Gateway    │
└─────────┬──────────┘
          │
          │ gRPC
          │ Protocol Buffers
          │
          ▼
┌────────────────────┐
│                    │
│   Go gRPC Server   │
│   localhost:50051  │
│                    │
└─────────┬──────────┘
          │
          ▼
┌────────────────────┐
│       Store        │
│                    │
│ map[int64]Todo     │
│                    │
└────────────────────┘