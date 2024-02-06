package ports

import (
	"cloud-api/backend/vnet-svc/app"
	"encoding/json"
)

type HttpServer struct {
	buildSVC app.VNetSVC
}

func NewHttpServer(buildSVC app.VNetSVC) HttpServer {
	return HttpServer{
		buildSVC: buildSVC,
	}
}

func (h HttpServer) ExchangeStructs(val interface{}, r interface{}) error {
	var err error

	byteData, err := json.Marshal(val)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, &r)
	// err = json.NewDecoder(stringBuffer).Decode(&r)
	if err != nil {
		return err
	}
	return nil
}
