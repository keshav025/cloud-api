package domain

type VnetCommand interface {
	CreateOrUpdateVnet(vnet VNet) error
	DeleteVNet(vnet VNet) error
}
