package command

import (
	"context"

	"cloud-api/backend/vnet-svc/domain"

	"errors"
)

type DeleteVNetHandler struct {
	VNetRepo      domain.Command
	VNetCloudRepo domain.VnetCommand
}

func NewDeleteVNetHandler(VNetRepo domain.Command) DeleteVNetHandler {
	if VNetRepo == nil {
		panic("nil VNetRepo")
	}

	return DeleteVNetHandler{VNetRepo: VNetRepo}
}

func (l DeleteVNetHandler) Handle(ctx context.Context, vnetId int32) (err error) {
	queryRepo := l.VNetRepo.(domain.Query)
	res := domain.VNet{}
	err = queryRepo.Get(ctx, &res, vnetId)
	if err != nil {
		return errors.New("unable-to-delete-VNet")
	}
	go l.DeleteVNet(ctx, res)
	return nil
}

func (l DeleteVNetHandler) DeleteVNet(ctx context.Context, vnet domain.VNet) (err error) {
	startStatus := "delete-in-progress"
	completedStatus := "deleted"
	vnet.Status = &startStatus
	err = l.VNetRepo.Update(ctx, &vnet, int32(vnet.ID))
	if err != nil {
		return errors.New("unable-to-delete-VNet")
	}
	l.VNetCloudRepo, err = GetAdapter(vnet.Platform)
	if err != nil {
		completedStatus = "failed"
		return errors.New("unable-to-create-VNet")
	}
	defer func() {
		vnet.Status = &completedStatus
		l.VNetRepo.Update(ctx, &vnet, int32(vnet.ID))

	}()

	err = l.VNetCloudRepo.DeleteVNet(vnet)
	if err != nil {
		completedStatus = "failed"
		return errors.New("unable-to-delete-VNet")
	}

	return nil
}
