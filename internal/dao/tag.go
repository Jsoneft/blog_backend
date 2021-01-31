package dao

import (
	"ginblog_backend/internal/model"
	"ginblog_backend/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Model: nil,
		Name:  name,
		State: state,
	}
	return tag.Count(d.engine)
}

func (d Dao) GetTag(tagID uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{
		Model: &model.Model{ID: tagID},
		State: state,
	}
	return tag.Get(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Model: nil,
		Name:  name,
		State: state,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			CreatedBy: createdBy,
		},
		Name:  name,
		State: state,
	}
	return tag.Create(d.engine)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	return tag.Delete(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name, updatedBy string, state uint8) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"state":      state,
		"updated_by": updatedBy,
	}
	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.engine, values)
}
