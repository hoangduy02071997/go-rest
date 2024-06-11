package business

import (
	"context"
	"rest-api/modules/item/model"
	"strings"
)

type UpdateTodoItemBusiness interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, data *model.TodoItemUpdate, condition map[string]interface{}) error
}

type updateTodoItemBusiness struct {
	store UpdateTodoItemBusiness
}

func NewUpdateTodoItemBusiness(store UpdateTodoItemBusiness) *updateTodoItemBusiness {
	return &updateTodoItemBusiness{store: store}
}

func (business *updateTodoItemBusiness) UpdateItemById(ctx context.Context, data *model.TodoItemUpdate, id int) error {
	instance, err := business.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if instance.Status != nil && *instance.Status == model.ItemStatuDeleted {
		return model.ErrItemDeleted
	}

	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrTitleEmpty
	}

	if err := business.store.UpdateItem(ctx, data, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
