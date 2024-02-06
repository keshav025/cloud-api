package azure

import (
	"cloud-api/backend/vnet-svc/config"
	"sync"
)

var vnetSvcRepo *VNetSvcRepository
var once sync.Once
var errCon error

type VNetSvcRepository struct {
	SubscriptionID string
	ResourceGroup  string
	ClientId       string
	TenantId       string
	ClientSecret   string
	l              *sync.Mutex
}

func NewVNetSvcRepository() (*VNetSvcRepository, error) {

	once.Do(func() {

		cfg := config.GetConfig()

		vnetSvcRepo = &VNetSvcRepository{
			SubscriptionID: cfg.AzureSubscriptionId,
			ResourceGroup:  cfg.AzureResourceGroup,
			ClientId:       cfg.AzureClientId,
			ClientSecret:   cfg.AzureClientSecret,
			TenantId:       cfg.AzureTenantId,
			l:              &sync.Mutex{},
		}
	})

	return vnetSvcRepo, errCon
}
