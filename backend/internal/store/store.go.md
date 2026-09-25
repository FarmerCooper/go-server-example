package store

import (
    "sync"
    "time"
)

type Todo struct {
    ID        int64
    Title     string
    Completed bool
    CreatedAt time.Time
}

type Store struct {
    mu     sync.RWMutex
    nextID int64
    todos  map[int64]Todo
}

func New() *Store {
    return &Store{
        nextID: 1,
        todos:  make(map[int64]Todo),
    }
}

func (s *Store) List() []Todo {
    s.mu.RLock()
    defer s.mu.RUnlock()

    todos := make([]Todo, 0, len(s.todos))

    for _, todo := range s.todos {
        todos = append(todos, todo)
    }

    return todos
}

func (s *Store) Create(title string, completed bool) Todo {
    s.mu.Lock()
    defer s.mu.Unlock()

    todo := Todo{
        ID:        s.nextID,
        Title:     title,
        Completed: completed,
        CreatedAt: time.Now().UTC(),
    }

    s.todos[todo.ID] = todo
    s.nextID++

    return todo
}

func (s *Store) Get(id int64) (Todo, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    todo, found := s.todos[id]

    return todo, found
}