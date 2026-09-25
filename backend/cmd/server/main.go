package main

import (
    "log"
    "net"
    "os"

    todov1 "example.com/todo-platform/backend/api/todo/v1"
    "example.com/todo-platform/backend/internal/database"
    appserver "example.com/todo-platform/backend/internal/server"
    "example.com/todo-platform/backend/internal/store"

    "google.golang.org/grpc"
)

func main() {
    dsn := os.Getenv("DATABASE_DSN")

    if dsn == "" {
        dsn = "host=localhost user=todo password=todo dbname=todo port=5432 sslmode=disable TimeZone=UTC"
    }

    db, err := database.Open(dsn)
    if err != nil {
        log.Fatal(err)
    }

    if err := db.AutoMigrate(
        &store.Todo{},
    ); err != nil {
        log.Fatal(err)
    }

    todoStore := store.New(db)

    todoServer := appserver.New(
        todoStore,
    )

    grpcServer := grpc.NewServer()

    todov1.RegisterTodoServiceServer(
        grpcServer,
        todoServer,
    )

    listener, err := net.Listen(
        "tcp",
        ":50051",
    )

    if err != nil {
        log.Fatal(err)
    }

    log.Println(
        "gRPC Todo server listening on :50051",
    )

    if err := grpcServer.Serve(
        listener,
    ); err != nil {
        log.Fatal(err)
    }
}