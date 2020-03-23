package ums

import (
	dalums "wehelp-api/dal/ums"
)

// AddressDTO : data transfer object  (address) table.
type AddressDTO struct {
	ID        int32  `json:"id"`
	Address   string `json:"address"`
	Latlng    string `json:"latlng"`
	City      string `json:"city"`
	IsActive  bool   `json:"is_active"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// AddressDTOToDAL : convert AddressDTO to AddressDAL
func (a *AddressDTO) AddressDTOToDAL() (*dalums.AddressDAL, error) {
	address := &dalums.AddressDAL{
		ID:        a.ID,
		Address:   a.Address,
		Latlng:    a.Latlng,
		City:      a.City,
		IsActive:  a.IsActive,
		CreatedAt: a.CreatedAt,
	}
	return address, nil
}

// AddressDALToDTO : convert AddressDAL to AddressDTO
func AddressDALToDTO(a *dalums.AddressDAL) (*AddressDTO, error) {
	address := &AddressDTO{
		ID:        a.ID,
		Address:   a.Address,
		Latlng:    a.Latlng,
		City:      a.City,
		IsActive:  a.IsActive,
		CreatedAt: a.CreatedAt,
	}
	return address, nil
}

// AddressDALToDTOArr : convert Array of AddressDAL to Array of AddressDTO
func AddressDALToDTOArr(addresses []*dalums.AddressDAL) ([]*AddressDTO, error) {
	var err error
	res := make([]*AddressDTO, len(addresses))
	for i, address := range addresses {
		res[i], err = AddressDALToDTO(address)
		if err != nil {
			return res, err
		}
	}
	return res, nil
}
