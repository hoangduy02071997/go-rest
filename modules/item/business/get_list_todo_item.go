package business

import (
	"context"
	"rest-api/common"
	"rest-api/modules/item/model"
)

type GetListItemStorage interface {
	GetAllItems(
		ctx context.Context,
		filter *model.FilterItem,
		paging *common.Paging,
		moreKey ...string,
	) ([]model.TodoItem, error)
}

type getListItemBusiness struct {
	store GetListItemStorage
}

func NewGetListItemBusiness(store GetListItemStorage) *getListItemBusiness {
	return &getListItemBusiness{store}
}

func (business *getListItemBusiness) GetListItems(
	ctx context.Context,
	filter *model.FilterItem,
	paging *common.Paging,
) ([]model.TodoItem, error) {
	data, err := business.store.GetAllItems(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
