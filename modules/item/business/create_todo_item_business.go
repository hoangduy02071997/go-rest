package business

import (
	"context"
	"rest-api/common"
	"rest-api/modules/item/model"
	"strings"
)

type CreateItemStorage interface {
	InsertItem(ctx context.Context, data *model.TodoItemCreation) error
}

type createTodoItemBusiness struct {
	store CreateItemStorage
}

func NewCreateTodoItemBusiness(store CreateItemStorage) *createTodoItemBusiness {
	return &createTodoItemBusiness{store: store}
}

func (business *createTodoItemBusiness) CreateNewTodoItemBusiness(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleEmpty
	}

	if data.Status == nil {
		dfStatus, _ := model.ParseItemStatus("todo")
		data.Status = &dfStatus
	}
	if err := business.store.InsertItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
