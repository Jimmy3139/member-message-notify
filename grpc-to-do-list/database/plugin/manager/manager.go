package manager

import (
	"gorm.io/gorm"
)

type ManagerPlugin struct{}

func (plugin *ManagerPlugin) Name() string {
	return "manager_plugin"
}
func (plugin *ManagerPlugin) Initialize(db *gorm.DB) error {
	db.Callback().Create().Before("gorm:create").Register("manager_plugin:before_create", plugin.BeforeCreate)
	return nil
}

func (*ManagerPlugin) BeforeCreate(db *gorm.DB) {
	if db.Statement.Schema != nil {
		ctx := db.Statement.Context
		for i := 0; i < len(db.Statement.Schema.Fields); i++ {
			if db.Statement.Schema.Fields[i].DBName == "created_by" {
				db.Statement.SetColumn("created_by", ctx.Value(ManagerPlugin{}).(int64))
			}
		}
	}
}
