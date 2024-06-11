package mysql

import (
	"context"
	"rest-api/modules/item/model"
)

func (s *sqlStorage) DeleteItemById(ctx context.Context, condition map[string]interface{}) error {
	deletedStatus := model.ItemStatuDeleted

	if err := s.db.Table(model.TodoItem{}.TableName()).Where(condition).Updates(
		map[string]interface{}{"status": deletedStatus.String()}).Error; err != nil {
		return err
	}
	return nil
}
