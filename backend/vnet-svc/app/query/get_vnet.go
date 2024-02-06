package query

import (
	"context"

	"cloud-api/backend/vnet-svc/domain"

	"errors"
)

type GetVNetHandler struct {
	VNetRepo domain.Query
}

func NewGetVNetHandler(VNetRepo domain.Query) GetVNetHandler {
	if VNetRepo == nil {
		panic("nil VNetRepo")
	}

	return GetVNetHandler{VNetRepo: VNetRepo}
}

func (l GetVNetHandler) Handle(ctx context.Context, vnetId int32) (res domain.VNet, err error) {

	err = l.VNetRepo.Get(ctx, &res, vnetId)
	if err != nil {
		return res, errors.New("unable-to-Get-VNet")
	}
	return res, nil
}
