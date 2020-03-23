package ums

import (
	"wehelp-api/db"
)

// AddressDAL : data access layer  (address) table.
type AddressDAL struct {
	ID        int32  `json:"id" gorm:"column:id;primary_key:true"`
	Address   string `json:"address" gorm:"column:address"`
	Latlng    string `json:"latlng" gorm:"column:latlng"`
	City      string `json:"city" gorm:"column:city"`
	IsActive  bool   `json:"is_active" gorm:"column:is_active"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64  `json:"updated_at" gorm:"column:updated_at"`
}

// TableName sets the insert table name for this struct type
func (a *AddressDAL) TableName() string {
	return "address"
}

// GetAllAddresses : get all addresses.
func GetAllAddresses() []*AddressDAL {
	addresses := []*AddressDAL{}
	db.DB().Find(&addresses)
	return addresses
}

// GetAddress : get one address by id.
func GetAddress(id int32) (*AddressDAL, error) {
	a := &AddressDAL{}
	result := db.DB().First(a, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}

// CreateAddress : create new address.
func CreateAddress(a *AddressDAL) (*AddressDAL, error) {
	result := db.DB().Create(a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}

// UpdateAddress : update exist address.
func UpdateAddress(a *AddressDAL) (*AddressDAL, error) {
	_, err := GetAddress(a.ID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}

// DeleteAddress : delete address by id.
func DeleteAddress(id int32) error {
	a, err := GetAddress(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(a)
	return result.Error
}
