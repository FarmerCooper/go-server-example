package store

import (
    "context"
    "errors"

    "gorm.io/gorm"
)

type Store struct {
    db *gorm.DB
}

func New(db *gorm.DB) *Store {
    return &Store{
        db: db,
    }
}

func (s *Store) List(
    ctx context.Context,
    completed *bool,
    priority *int64,
) ([]Todo, error) {
    var todos []Todo

    query := s.db.
        WithContext(ctx).
        Order("id ASC")

    if completed != nil {
        query = query.Where(
            "completed = ?",
            *completed,
        )
    }

    if priority != nil {
        query = query.Where(
            "priority = ?",
            *priority,
        )
    }

    if err := query.Find(&todos).Error; err != nil {
        return nil, err
    }

    return todos, nil
}

func (s *Store) Create(
    ctx context.Context,
    title string,
    completed bool,
    priority int64,
) (Todo, error) {
    todo := Todo{
        Title:     title,
        Completed: completed,
        Priority: priority,
    }

    if err := s.db.
        WithContext(ctx).
        Create(&todo).
        Error; err != nil {

        return Todo{}, err
    }

    return todo, nil
}

func (s *Store) Get(
    ctx context.Context,
    id int64,
) (Todo, bool, error) {
    var todo Todo

    err := s.db.
        WithContext(ctx).
        First(&todo, id).
        Error

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return Todo{}, false, nil
    }

    if err != nil {
        return Todo{}, false, err
    }

    return todo, true, nil
}

func (s *Store) Update(
    ctx context.Context,
    id int64,
    title *string,
    completed *bool,
    priority *int64,
) (Todo, bool, int64, error) {
    var todo Todo

    err := s.db.
        WithContext(ctx).
        First(&todo, id).
        Error

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return Todo{}, false, 0, nil
    }

    if err != nil {
        return Todo{}, false, 0, err
    }

    updates := map[string]any{}

    if title != nil {
        updates["title"] = *title
    }

    if completed != nil {
        updates["completed"] = *completed
    }

    if priority != nil {
        updates["priority"] = *priority
    }

    if len(updates) > 0 {
        if err := s.db.
            WithContext(ctx).
            Model(&todo).
            Updates(updates).
            Error; err != nil {

            return Todo{}, false, 0, err
        }
    }

    if err := s.db.
        WithContext(ctx).
        First(&todo, id).
        Error; err != nil {

        return Todo{}, false, 0, err
    }

    return todo, true, 0, nil
}

func (s *Store) Delete(
    ctx context.Context,
    id int64,
) (bool, error) {
    result := s.db.
        WithContext(ctx).
        Delete(&Todo{}, id)

    if result.Error != nil {
        return false, result.Error
    }

    return result.RowsAffected > 0, nil
}