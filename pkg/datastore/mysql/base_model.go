package mysql

import (
	"com/chat/service/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint64         `gorm:"column:id;primaryKey;" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

var s, _ = utils.NewSnowflake(1)

func (baseModel *BaseModel) BeforeCreate(_ *gorm.DB) (err error) {
	generate, err := s.Generate()

	baseModel.ID = uint64(generate)
	return nil
}
