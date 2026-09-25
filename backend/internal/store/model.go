package store

import "time"

type Todo struct {
    ID int64 `gorm:"primaryKey;autoIncrement"`

    Title string `gorm:"size:200;not null"`

    Completed bool `gorm:"not null;default:false"`

    Priority int64 `gorm:"not null;default:3"`

    DueDate time.Time

    CreatedAt time.Time

    UpdatedAt time.Time
}