package ums

import (
	dalums "wehelp-api/dal/ums"
	"wehelp-api/utils"
)

// UserDTO : data transfer object  (user) table.
type UserDTO struct {
	ID         int32      `json:"id"`
	Email      string     `json:"email"`
	Name       string     `json:"name"`
	Password   string     `json:"password"`
	Phone      string     `json:"phone"`
	Bio        string     `json:"bio"`
	AddressDTO AddressDTO `json:"user_address"`
	IsActive   bool       `json:"is_active"`
	CreatedAt  int64      `json:"created_at"`
	UpdatedAt  int64      `json:"updated_at"`
}

// UserDTOToDAL : convert UserDTO to UserDAL
func (a *UserDTO) UserDTOToDAL() (*dalums.UserDAL, error) {
	user := &dalums.UserDAL{
		ID:          a.ID,
		Email:       a.Email,
		Name:        a.Name,
		PasswordMd5: utils.StringToMd5(a.Password),
		Phone:       utils.SetNullString(a.Phone),
		Bio:         utils.SetNullString(a.Bio),
		IsActive:    a.IsActive,
	}
	// TODO: check user address
	// user.Address =     a.AddressDTO.ID,

	return user, nil
}

// UserDALToDTO : convert UserDAL to UserDTO
func UserDALToDTO(a *dalums.UserDAL) (*UserDTO, error) {
	user := &UserDTO{
		ID:        a.ID,
		Email:     a.Email,
		Name:      a.Name,
		Phone:     a.Phone.String,
		Bio:       a.Bio.String,
		IsActive:  a.IsActive,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}

	// TODO: check user address
	// user.Address =     a.AddressDTO.ID,

	return user, nil
}

// UserDALToDTOArr : convert Array of UserDAL to Array of UserDTO
func UserDALToDTOArr(useres []*dalums.UserDAL) ([]*UserDTO, error) {
	var err error
	res := make([]*UserDTO, len(useres))
	for i, user := range useres {
		res[i], err = UserDALToDTO(user)
		if err != nil {
			return res, err
		}
	}
	return res, nil
}
