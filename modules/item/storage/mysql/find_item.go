package mysql

import (
	"context"
	"gorm.io/gorm"
	"rest-api/common"
	"rest-api/modules/item/model"
)

func (s *sqlStorage) GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error) {
	var item model.TodoItem
	if err := s.db.Where(condition).First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(model.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &item, nil
}
