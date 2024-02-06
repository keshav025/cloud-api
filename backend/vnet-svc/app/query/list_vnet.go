package query

import (
	"context"
	"strings"

	"cloud-api/backend/vnet-svc/domain"

	"errors"
)

type ListVNetHandler struct {
	VNetRepo domain.Query
}

func NewListVNetHandler(VNetRepo domain.Query) ListVNetHandler {
	if VNetRepo == nil {
		panic("nil VNetRepo")
	}

	return ListVNetHandler{VNetRepo: VNetRepo}
}

func (l ListVNetHandler) Handle(ctx context.Context, platform, status *string) (res []domain.VNet, err error) {

	query := ""
	queryParams := []interface{}{}
	queries := []interface{}{}

	if status != nil {
		query = query + "status IN ?"
		queryParams = append(queryParams, strings.Split(*status, ","))
	}
	if platform != nil {
		if len(queryParams) > 0 {
			query = query + " AND "
		}
		query = query + "platform = ?"

		queryParams = append(queryParams, platform)

	}

	queries = append(queries, query)
	for _, q := range queryParams {
		queries = append(queries, q)
	}

	err = l.VNetRepo.List(ctx, &res, queries...)
	if err != nil {
		return res, errors.New("unable-to-List-VNet")
	}
	return res, nil
}
