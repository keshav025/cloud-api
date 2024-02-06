package domain

import "gorm.io/gorm"

type VNet struct {
	gorm.Model
	Name     string
	Location string `json:"location,omitempty"`
	CIDR     string `gorm:"column:cidr"`
	Platform string
	Status   *string `json:"status,omitempty"`
}

func (VNet) TableName() string {
	return "virtual_networks"
}
