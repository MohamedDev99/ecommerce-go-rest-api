package types

import "github.com/MohamedDev99/ecommerce-go-rest-api/pkg/db/models"

type AddressRepository interface {
	AddAddress(address models.Address) error
	// GetAddress(addressId string) (models.Address, error)
	// GetAllAddress() ([]models.Address, error)
	DeleteAddress(addressId string) error
	EditHomeAddress(addressId, homeAddress string) error
	EditWorkAddress(addressId, workAddress string) error
}
