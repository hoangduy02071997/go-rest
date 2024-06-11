package mysql

import (
	"context"
	"rest-api/common"
	"rest-api/modules/item/model"
)

func (s *sqlStorage) InsertItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		//return err // -> default root err in gorm
		return common.ErrDB(err) // ~~> use wrap root error to *AppError
	}
	return nil
}
