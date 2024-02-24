package database

import (
	"context"
	"grpc-to-do-list/database/plugin/manager"
	"grpc-to-do-list/database/sysconst"
	"grpc-to-do-list/database/utils"
	"grpc-to-do-list/dtos"
	"math"
	"reflect"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Base-Repo
type BaseRepo struct {
	*gorm.DB
}
type IBaseRepo interface {
	FindAndCount(e IEntity, pagination *utils.Pagination, condition ...func(db *gorm.DB) *gorm.DB) []map[string]interface{}
	FindByID(e IEntity) (state int)
	Create(createrId int64, e IEntity) (state int)
	Update(updaterId int64, e IEntity) (state int)
	Deletes(deleterId int64, e IEntity, ids []int64) (state int)
}

func (sql *BaseRepo) FindAndCount(e IEntity, p2 *dtos.Pagination, condition ...func(db *gorm.DB) *gorm.DB) []map[string]interface{} {
	var totalRows int64
	sql.DB.Table(e.TableName()).Scopes(condition...).Count(&totalRows)

	pagination := &utils.Pagination{
		Limit: p2.Limit,
		Page:  p2.Page,
		Sort:  p2.Sort,
	}

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	var results []map[string]interface{}
	sql.DB.Debug().
		Table(e.TableName()).           // SELECT
		Scopes(condition...).           // WHERE
		Limit(pagination.GetLimit()).   // LIMIT
		Offset(pagination.GetOffset()). // OFFSET
		Order(pagination.GetSort()).    // ORDER BY
		Find(&results)
	return results
}
func (sql *BaseRepo) FindByID(e IEntity) (state int) {
	result := sql.DB.Table(e.TableName()).Where("id = ?", e.GetID()).Find(&e)
	if result.RowsAffected <= 0 {

		return sysconst.InvalidId
	}
	return sysconst.Success
}
func (sql *BaseRepo) Create(createrId int64, e IEntity) (state int) {
	ctx := context.WithValue(context.Background(), manager.ManagerPlugin{}, createrId)
	tx := sql.DB.WithContext(ctx)
	if err := tx.Table(e.TableName()).Create(e).Error; err != nil {
		return sysconst.InsertError
	}
	return sysconst.Success
}
func (sql *BaseRepo) Update(updaterId int64, e IEntity) (state int, err error) {
	tx := sql.DB.Begin() // 需要查詢是否已經 begin()
	result := map[string]interface{}{}
	if err := tx.Table(e.TableName()).
		Where("id = ?", e.GetID()).
		Find(&result).Error; err != nil {
		return sysconst.InvalidId, err
	}

	updateData := removeNilFields(reflect.ValueOf(e).Elem())
	updateData["updated_by"] = updaterId
	updateData["updated_at"] = time.Now()

	if err := tx.Table(e.TableName()).Debug().
		Where("id = ?", e.GetID()).
		Updates(updateData).Error; err != nil {
		tx.Rollback()
		return sysconst.UpdateError, err
	}
	tx.Commit()
	return sysconst.Success, nil
}
func (sql *BaseRepo) Deletes(deleterId int64, e IEntity, ids []int64) (state int) {
	tx := sql.DB.Begin()
	if len(ids) == 0 {
		return sysconst.DelError
	}
	if err := tx.Table(e.TableName()).
		Where("id IN ? AND deleted_at IS NULL AND is_del = 0 AND deleted_by IS NULL", ids).
		Updates(map[string]interface{}{
			"deleted_by": deleterId,
			"deleted_at": time.Now(),
			"is_del":     1,
		}).Error; err != nil {
		tx.Rollback()
		return sysconst.DelError
	}

	tx.Commit()
	return sysconst.Success
}
func removeNilFields(v reflect.Value) map[string]interface{} {
	result := map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fType := v.Type().Field(i)

		// 允許 database.Entity 使用，且 排除 Join 相關的 Struct
		if f.Kind() == reflect.Struct && fType.Name == "Entity" {
			removeNilFields(f)
		}

		if !f.IsZero() && f.Kind() != reflect.Struct {
			column := schema.ParseTagSetting(fType.Tag.Get("gorm"), ";")["COLUMN"]
			result[column] = f.Interface()
		}
	}
	return result
}

// func (sql *BaseRepo) Eq(e IEntity) *gorm.DB {
// 	var scopes []func(db *gorm.DB) *gorm.DB
// 	v := reflect.ValueOf(e).Elem()
// 	for i := 0; i < v.NumField(); i++ {
// 		if !v.Field(i).IsZero() {

// 			fieldType := v.Type().Field(i)
// 			fieldVal := v.Field(i)

// 			if fieldVal.Kind() == reflect.Ptr {
// 				fieldVal = fieldVal.Elem()
// 			}
// 			scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
// 				column := schema.ParseTagSetting(fieldType.Tag.Get("gorm"), ";")["COLUMN"]
// 				return db.Where(e.TableName()+"."+column+" = ?", fieldVal.Interface())
// 			})
// 		}
// 	}
// 	return sql.DB.Debug().Table(e.TableName()).Scopes(scopes...)
// }
