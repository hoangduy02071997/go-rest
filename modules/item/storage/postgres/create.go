package postgres

import (
	"context"
	"rest-api/modules/item/model"
)

func (s *sqlStorage) InsertItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
