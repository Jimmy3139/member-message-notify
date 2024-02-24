package database

import (
	"time"
)

// Base-Entity
type Entity struct {
	ID        int64      `gorm:"column:id;type:bigint;primaryKey" mapstructure:"id"`
	CreatedBy *int64     `gorm:"column:created_by;type:bigint;" mapstructure:"created_by"`
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;DEFAULT NULL" mapstructure:"created_at"`
	UpdatedBy *int64     `gorm:"column:updated_by;type:bigint;" mapstructure:"updated_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;DEFAULT NULL" mapstructure:"updated_at"`
	DeletedBy *int64     `gorm:"column:deleted_by;type:bigint;" mapstructure:"deleted_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime;DEFAULT NULL" mapstructure:"deleted_at"`
	IsDel     *bool      `gorm:"column:is_del;type:tinyint;" mapstructure:"is_del"`
}

type IEntity interface {
	TableName() string
	GetID() int64
}

func (entity *Entity) GetID() int64 {
	return entity.ID
}
