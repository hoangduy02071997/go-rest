package postgres

import (
	"context"
	"rest-api/common"
	"rest-api/modules/item/model"
)

func (s *sqlStorage) GetAllItems(
	ctx context.Context,
	filter *model.FilterItem,
	paging *common.Paging,
	moreKey ...string,
) ([]model.TodoItem, error) {
	db := s.db.Where("status != ?", "deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	var todoItems []model.TodoItem
	if err := db.Debug().Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Debug().Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&todoItems).
		Error; err != nil {
		return nil, err
	}

	return todoItems, nil
}
