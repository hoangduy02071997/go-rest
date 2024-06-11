package model

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"id"`
	Title       string      `json:"title" gorm:"column:titles"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (t TodoItemCreation) TableName() string { return "todos" }
