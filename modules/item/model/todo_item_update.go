package model

type TodoItemUpdate struct {
	Title       string      `json:"title" gorm:"column:title"`
	Description *string     `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (t TodoItemUpdate) TableName() string { return "todos" }
