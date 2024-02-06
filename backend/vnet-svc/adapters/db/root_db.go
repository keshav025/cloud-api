package db

import (
	"context"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r VNetDBSvcRepository) Create(ctx context.Context, obj interface{}) error {
	res := r.DB.Create(obj)
	return res.Error

}
func (r VNetDBSvcRepository) Update(ctx context.Context, obj interface{}, objId int32) error {
	rval := reflect.ValueOf(obj)

	objCurrent := reflect.New(rval.Elem().Type()).Interface()
	r.l.Lock()
	defer r.l.Unlock()
	tx := r.DB.Begin()
	res := tx.First(&objCurrent, objId)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	rvalCurrent := reflect.ValueOf(objCurrent)
	irvalCurr := reflect.Indirect(rvalCurrent)
	irValType := irvalCurr.Type()

	for i, limit := 0, irValType.NumField(); i < limit; i++ {
		fld := irValType.Field(i)

		v := irvalCurr.FieldByName(fld.Name).Addr().Interface()
		model, ok := v.(*gorm.Model)
		if ok {
			s := rval.Elem()
			if s.Kind() == reflect.Struct {
				// exported field
				f := s.FieldByName("Model")
				if f.IsValid() {

					if f.CanSet() {

						if f.Kind() == reflect.Struct {
							f.FieldByName("CreatedAt").Set(reflect.ValueOf(model.CreatedAt))

						}
					}
				}
			}
			break
		}
	}

	res = tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(obj)

	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		// handle error
		return err
	}

	return res.Error

}

func (r VNetDBSvcRepository) Delete(ctx context.Context, obj interface{}, objId int32) error {
	res := r.DB.Delete(obj, objId)
	return res.Error
}

func (r VNetDBSvcRepository) Get(ctx context.Context, obj interface{}, objId int32) error {

	res := r.DB.
		Preload(clause.Associations, func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		}).
		First(obj, objId)
	return res.Error

}

func (r VNetDBSvcRepository) List(ctx context.Context, objList interface{}, filters ...interface{}) error {

	res := r.DB.
		Preload(clause.Associations, func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		}).
		Order("created_at DESC").Find(objList, filters...)
	return res.Error
}
