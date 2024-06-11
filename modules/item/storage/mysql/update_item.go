package mysql

import (
	"context"
	"rest-api/modules/item/model"
)

func (s *sqlStorage) UpdateItem(ctx context.Context, data *model.TodoItemUpdate, condition map[string]interface{}) error {
	if err := s.db.Where(condition).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
