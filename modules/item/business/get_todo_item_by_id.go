package business

import (
	"context"
	"rest-api/modules/item/model"
)

type GetTodoItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
}

type getTodoItemBusiness struct {
	store GetTodoItemStorage
}

func NewGetTodoItemBusiness(store GetTodoItemStorage) *getTodoItemBusiness {
	return &getTodoItemBusiness{store: store}
}

func (business *getTodoItemBusiness) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := business.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
