package model

import (
	"errors"
	"rest-api/common"
)

const (
	EntityName = "todo_item"
)

var (
	ErrTitleEmpty  = errors.New("title cannot be empty")
	ErrItemDeleted = errors.New("item was deleted")
)

type TodoItem struct {
	Title           string      `json:"title" gorm:"title"`
	Description     string      `json:"description" gorm:"description"`
	Status          *ItemStatus `json:"status" gorm:"status"`
	common.SQLModel             // Sử dũng ký thuật embedding (lấy toàn bộ struct ra struct này) ~~ Embed struct
}

func (t TodoItem) TableName() string { return "todos" }
