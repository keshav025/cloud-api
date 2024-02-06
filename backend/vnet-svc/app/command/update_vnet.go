package command

import (
	"context"

	"cloud-api/backend/vnet-svc/domain"
)

type UpdateVNetHandler struct {
	CreateVNetHandler
}

func NewUpdateVNetHandler(VNetRepo CreateVNetHandler) UpdateVNetHandler {

	return UpdateVNetHandler{CreateVNetHandler: VNetRepo}
}

func (l UpdateVNetHandler) Handle(ctx context.Context, VNet domain.VNet, vnetId int32) (err error) {

	go l.CreateVNetHandler.CreateOrUpdateVnets(ctx, VNet)
	return nil
}
