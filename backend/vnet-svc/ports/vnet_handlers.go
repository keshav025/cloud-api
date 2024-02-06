package ports

import (
	"cloud-api/backend/common/server/httperr"
	"cloud-api/backend/vnet-svc/domain"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

func (h HttpServer) CreateVNet(w http.ResponseWriter, r *http.Request) {

	vnet := VNet{}
	if err := render.Decode(r, &vnet); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	obj := domain.VNet{}
	err := h.ExchangeStructs(vnet, &obj)
	if err != nil {
		logrus.Errorf("error parsing vnet: %v", err)
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	err = h.buildSVC.Commands.CreateVNet.Handle(r.Context(), obj)
	if err != nil {
		logrus.Errorf("error creating vnet: %v", err)
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (h HttpServer) RetryVNet(w http.ResponseWriter, r *http.Request, vnetID int32) {

	// vnet := VNet{}
	// if err := render.Decode(r, &vnet); err != nil {
	// 	httperr.BadRequest("invalid-request", err, w, r)
	// 	return
	// }
	vnet, err := h.buildSVC.Queries.GetVNet.Handle(r.Context(), vnetID)
	if err != nil {
		logrus.Errorf("error getting vnet: %v", err)
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	err = h.buildSVC.Commands.UpdateVNet.Handle(r.Context(), vnet, vnetID)
	if err != nil {
		logrus.Errorf("error updating vnet: %v", err)
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
func (h HttpServer) DeleteVNet(w http.ResponseWriter, r *http.Request, vnetID int32) {

	err := h.buildSVC.Commands.DeleteVNet.Handle(r.Context(), vnetID)
	if err != nil {
		logrus.Errorf("error deleting vnet: %v", err)
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
func (h HttpServer) GetVNet(w http.ResponseWriter, r *http.Request, vnetID int32) {

	vnet, err := h.buildSVC.Queries.GetVNet.Handle(r.Context(), vnetID)
	if err != nil {
		logrus.Errorf("error getting vnet: %v", err)
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	vnetRes := VNet{}

	err = h.ExchangeStructs(vnet, &vnetRes)
	if err != nil {
		logrus.Errorf("error parsing vnet: %v", err)
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	render.Respond(w, r, vnetRes)

}

func (h HttpServer) ListVNet(w http.ResponseWriter, r *http.Request, params ListVNetParams) {

	vnetList, err := h.buildSVC.Queries.ListVNet.Handle(r.Context(), params.Platform, params.Status)
	if err != nil {
		logrus.Errorf("error retrieving vnet list : %v", err)
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	vnetListRes := []VNet{}

	err = h.ExchangeStructs(vnetList, &vnetListRes)
	if err != nil {
		logrus.Errorf("error parsing vnet: %v", err)
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	// for _, vnet := range vnetList {

	// 	// id := int32(vnet.Model.ID)
	// 	vnetRes := h.GetVNetDTO(vnet)
	// 	vnetListRes = append(vnetListRes, *vnetRes)

	// }
	render.Respond(w, r, vnetListRes)

}
