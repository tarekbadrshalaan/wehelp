package ums

import (
	dalums "wehelp-api/dal/ums"
	dtoums "wehelp-api/dto/ums"
)

// GetAllAddresses : get All addresses.
func GetAllAddresses() ([]*dtoums.AddressDTO, error) {
	addresses := dalums.GetAllAddresses()
	return dtoums.AddressDALToDTOArr(addresses)
}

// GetAddress : get one address by id.
func GetAddress(id int32) (*dtoums.AddressDTO, error) {
	a, err := dalums.GetAddress(id)
	if err != nil {
		return nil, err
	}
	return dtoums.AddressDALToDTO(a)
}

// CreateAddress : create new address.
func CreateAddress(a *dtoums.AddressDTO) (*dtoums.AddressDTO, error) {
	address, err := a.AddressDTOToDAL()
	if err != nil {
		return nil, err
	}
	newaddress, err := dalums.CreateAddress(address)
	if err != nil {
		return nil, err
	}
	return dtoums.AddressDALToDTO(newaddress)
}

// UpdateAddress : update exist address.
func UpdateAddress(a *dtoums.AddressDTO) (*dtoums.AddressDTO, error) {
	address, err := a.AddressDTOToDAL()
	if err != nil {
		return nil, err
	}
	updateaddress, err := dalums.UpdateAddress(address)
	if err != nil {
		return nil, err
	}
	return dtoums.AddressDALToDTO(updateaddress)
}

// DeleteAddress : delete address by id.
func DeleteAddress(id int32) error {
	return dalums.DeleteAddress(id)
}
