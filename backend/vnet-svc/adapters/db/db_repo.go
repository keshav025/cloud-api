package db

import (
	"cloud-api/backend/vnet-svc/config"
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var vnetSvcRepo *VNetDBSvcRepository
var once sync.Once
var errCon error

type VNetDBSvcRepository struct {
	DB *gorm.DB
	l  *sync.Mutex
}

func NewVNetDBSvcRepository() (*VNetDBSvcRepository, error) {

	once.Do(func() {
		var db *gorm.DB
		cfg := config.GetConfig()
		dsn := fmt.Sprintf("host=%s port=%v user=%s "+
			// "password=%s dbname=%s sslmode=disable",
			"password=%s dbname=%s sslmode=require",
			cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

		db, errCon := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if errCon != nil {
			return
		}

		vnetSvcRepo = &VNetDBSvcRepository{
			DB: db,
			l:  &sync.Mutex{},
		}
	})

	return vnetSvcRepo, errCon
}
