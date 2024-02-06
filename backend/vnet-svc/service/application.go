package service

import (
	"context"
	"sync"

	dbAdapter "cloud-api/backend/vnet-svc/adapters/db"
	"cloud-api/backend/vnet-svc/app"
	"cloud-api/backend/vnet-svc/app/command"
	"cloud-api/backend/vnet-svc/app/query"
	"cloud-api/backend/vnet-svc/domain"
)

var appSvc app.VNetSVC
var once sync.Once
var err error

func NewApplication(ctx context.Context) app.VNetSVC {
	once.Do(func() {

		vnetSVCRepository, err := dbAdapter.NewVNetDBSvcRepository()
		if err != nil {
			panic(err)
		}

		masterDBConn := vnetSVCRepository.DB

		var models = []interface{}{&domain.VNet{}}

		masterDBConn.AutoMigrate(models...)
		appSvc = app.VNetSVC{
			Commands: app.Commands{
				CreateVNet: command.NewCreateVNetHandler(vnetSVCRepository),
				UpdateVNet: command.NewUpdateVNetHandler(command.NewCreateVNetHandler(vnetSVCRepository)),
				DeleteVNet: command.NewDeleteVNetHandler(vnetSVCRepository),
			},
			Queries: app.Queries{

				GetVNet:  query.NewGetVNetHandler(vnetSVCRepository),
				ListVNet: query.NewListVNetHandler(vnetSVCRepository),
			},
		}
	})
	return appSvc
}
