package app

import (
	"cloud-api/backend/vnet-svc/app/command"
	"cloud-api/backend/vnet-svc/app/query"
)

type VNetSVC struct {
	// CommandChan chan command.CommandHandler
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateVNet command.CreateVNetHandler
	UpdateVNet command.UpdateVNetHandler
	DeleteVNet command.DeleteVNetHandler
}

type Queries struct {
	GetVNet  query.GetVNetHandler
	ListVNet query.ListVNetHandler
}
