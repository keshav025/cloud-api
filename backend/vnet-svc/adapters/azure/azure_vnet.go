package azure

import (
	"cloud-api/backend/vnet-svc/domain"
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/network/mgmt/network"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/sirupsen/logrus"
)

func (a VNetSvcRepository) CreateOrUpdateVnet(vnet domain.VNet) error {

	cfg := auth.NewClientCredentialsConfig(
		a.ClientId,
		a.ClientSecret,
		a.TenantId,
	)

	authorizer, err := cfg.Authorizer()
	if err != nil {
		logrus.Error(err)
		return err
	}
	subscriptionID := a.SubscriptionID
	resourceGroup := a.ResourceGroup
	location := vnet.Location
	cidr := vnet.CIDR
	vnetName := vnet.Name
	// Create resource group
	groupsClient := resources.NewGroupsClient(subscriptionID)
	groupsClient.Authorizer = authorizer

	_, err = groupsClient.CreateOrUpdate(context.Background(), resourceGroup, resources.Group{
		Location: to.StringPtr(location),
	})
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Create VNet
	vnetClient := network.NewVirtualNetworksClient(subscriptionID)
	vnetClient.Authorizer = authorizer

	future, err := vnetClient.CreateOrUpdate(
		context.Background(),
		resourceGroup,
		vnetName,
		network.VirtualNetwork{
			Location: to.StringPtr(location),
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{cidr},
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Wait for completion
	err = future.WaitForCompletionRef(context.Background(), vnetClient.Client)

	return err
}

func (a VNetSvcRepository) DeleteVNet(vnet domain.VNet) error {

	cfg := auth.NewClientCredentialsConfig(
		a.ClientId,
		a.ClientSecret,
		a.TenantId,
	)

	authorizer, err := cfg.Authorizer()
	if err != nil {
		logrus.Error(err)
		return err
	}
	subscriptionID := a.SubscriptionID
	resourceGroup := a.ResourceGroup

	vnetName := vnet.Name

	// Create VNet
	vnetClient := network.NewVirtualNetworksClient(subscriptionID)
	vnetClient.Authorizer = authorizer

	future, err := vnetClient.Delete(context.Background(), resourceGroup, vnetName)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Wait for completion
	err = future.WaitForCompletionRef(context.Background(), vnetClient.Client)

	return err
}
