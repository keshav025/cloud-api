package command

import (
	"context"
	"strings"

	azureClient "cloud-api/backend/vnet-svc/adapters/azure"
	"cloud-api/backend/vnet-svc/domain"

	"errors"
)

type CreateVNetHandler struct {
	VNetRepo      domain.Command
	VneCloudtRepo domain.VnetCommand
}

func NewCreateVNetHandler(VNetRepo domain.Command) CreateVNetHandler {
	if VNetRepo == nil {
		panic("nil VNetRepo")
	}

	return CreateVNetHandler{VNetRepo: VNetRepo}
}

func (l CreateVNetHandler) Handle(ctx context.Context, VNet domain.VNet) (err error) {
	queryRepo := l.VNetRepo.(domain.Query)
	res := []domain.VNet{}
	err = queryRepo.List(ctx, &res, "name = ? AND cidr = ?", VNet.Name, VNet.CIDR)
	if err != nil {
		return errors.New("unable-to-create-VNet")
	}

	if len(res) > 0 {
		return errors.New("unable-to-create-VNet: duplicate record")
	}

	err = l.VNetRepo.Create(ctx, &VNet)
	if err != nil {
		return errors.New("unable-to-create-VNet")
	}
	go l.CreateOrUpdateVnets(ctx, VNet)
	return nil
}

func (l CreateVNetHandler) CreateOrUpdateVnets(ctx context.Context, vnet domain.VNet) (err error) {
	startStatus := "create-in-progress"
	completedStatus := "completed"
	vnet.Status = &startStatus
	err = l.VNetRepo.Update(ctx, &vnet, int32(vnet.ID))
	if err != nil {
		return errors.New("unable-to-create-VNet")
	}

	defer func() {
		vnet.Status = &completedStatus
		l.VNetRepo.Update(ctx, &vnet, int32(vnet.ID))

	}()
	l.VneCloudtRepo, err = GetAdapter(vnet.Platform)
	if err != nil {
		completedStatus = "failed"
		return errors.New("unable-to-create-VNet")
	}
	err = l.VneCloudtRepo.CreateOrUpdateVnet(vnet)
	if err != nil {
		completedStatus = "failed"
		return errors.New("unable-to-create-VNet")
	}

	return nil
}

func GetAdapter(platform string) (domain.VnetCommand, error) {
	if strings.ToLower(platform) == "azure" {
		return azureClient.NewVNetSvcRepository()
	}
	return nil, nil
}
