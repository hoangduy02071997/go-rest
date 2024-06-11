package business

import (
	"context"
	"rest-api/modules/item/model"
)

type DeleteTodoItemBusiness interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
	DeleteItemById(ctx context.Context, condition map[string]interface{}) error
}

type deleteTodoItemBusiness struct {
	store DeleteTodoItemBusiness
}

func NewDeleteTodoItemBusiness(store DeleteTodoItemBusiness) *deleteTodoItemBusiness {
	return &deleteTodoItemBusiness{store: store}
}

func (business *deleteTodoItemBusiness) DeleteItem(ctx context.Context, id int) error {
	instance, err := business.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if instance.Status != nil && *instance.Status == model.ItemStatuDeleted {
		return model.ErrItemDeleted
	}

	if err := business.store.DeleteItemById(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
