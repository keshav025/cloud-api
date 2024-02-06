// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package ports

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// VNet defines model for VNet.
type VNet struct {
	Cidr      *string `json:"cidr,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Id        *int32  `json:"id,omitempty"`
	Location  *string `json:"location,omitempty"`
	Name      *string `json:"name,omitempty"`
	Platform  *string `json:"platform,omitempty"`
	Status    *string `json:"status,omitempty"`
	VnetId    *int32  `json:"vnetId,omitempty"`
}

// ListVNetParams defines parameters for ListVNet.
type ListVNetParams struct {
	// Platform todo
	Platform *string `form:"platform,omitempty" json:"platform,omitempty"`

	// Status todo
	Status *string `form:"status,omitempty" json:"status,omitempty"`
}

// CreateVNetJSONRequestBody defines body for CreateVNet for application/json ContentType.
type CreateVNetJSONRequestBody = VNet