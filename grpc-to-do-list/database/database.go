package database

import (
	"fmt"
	//"grpc-to-do-list/database/conf"
	"grpc-to-do-list/config"
	"grpc-to-do-list/database/plugin/manager"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type DataBase struct {
	MySQL *BaseRepo
}

func initDB(masterDSN string, slaveDSNs []string) *gorm.DB {
	masterDialector := mysql.Open(masterDSN)
	slaveDialectors := []gorm.Dialector{}

	for _, dsn := range slaveDSNs {
		slaveDialector := mysql.Open(dsn)
		slaveDialectors = append(slaveDialectors, slaveDialector)
	}

	db, _ := gorm.Open(masterDialector)

	// 設定讀寫分離
	// gorm-dbresolver: https://github.com/go-gorm/dbresolver
	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{masterDialector},
		Replicas: slaveDialectors,
		Policy:   dbresolver.RandomPolicy{}, // 目前只有 RandomPolicy 這個選項
	}))
	// 設定 created_by / updated_by / deleted_by
	db.Use(&manager.ManagerPlugin{})

	// 設定 MySQL 連線狀態
	mysqlDB, _ := db.DB()

	mysqlDB.SetConnMaxLifetime(time.Minute * 3) // 客戶端主動斷開之分鐘數，MySQL 建議 5分內
	mysqlDB.SetMaxOpenConns(10)                 // 取決MySQL配置數量
	mysqlDB.SetMaxIdleConns(10)                 // MySQL 建議與SetMaxOpenConns相同

	return db
}

func convToDNS(c config.DBCollentSetting) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Name)
}

func New(conf config.DBCollentSetting) *DataBase {
	return &DataBase{
		MySQL: &BaseRepo{
			DB: initDB(
				convToDNS(conf),
				[]string{
					convToDNS(conf),
					// convToDNS(conf.Database.MySQL.Slave), // 目前 ReadOnly，同步延遲所以不使用
				},
			),
		},
	}
}
