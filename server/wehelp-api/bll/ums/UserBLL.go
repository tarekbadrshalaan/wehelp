package ums

import (
	"errors"
	"fmt"
	"time"
	dalums "wehelp-api/dal/ums"
	dtoums "wehelp-api/dto/ums"
	"wehelp-api/utils"

	"github.com/jinzhu/gorm"
)

// GetAllUseres : get All useres.
func GetAllUseres() ([]*dtoums.UserDTO, error) {
	useres := dalums.GetAllUseres()
	return dtoums.UserDALToDTOArr(useres)
}

// GetUser : get one user by id.
func GetUser(id int32) (*dtoums.UserDTO, error) {
	u, err := dalums.GetUser(id)
	if err != nil {
		return nil, err
	}
	return dtoums.UserDALToDTO(u)
}

// GetUserByEmail : get one user by email.
func GetUserByEmail(email string) (*dtoums.UserDTO, error) {
	u, err := dalums.GetUserByEmail(email)
	if err != nil {
		return nil, err //fmt.Errorf("errors founded %w", err)
	}
	return dtoums.UserDALToDTO(u)
}

// CreateUser : create new user.
func CreateUser(u *dtoums.UserDTO) (*dtoums.UserDTO, error) {
	// check if email is valid
	if valid := utils.ValidEmail(u.Email); !valid {
		return nil, fmt.Errorf("Not Valid Email (%v) ", u.Email)
	}

	// check if email already exists
	userEmail, err := GetUserByEmail(u.Email)
	if userEmail != nil {
		return nil, fmt.Errorf("Email (%v) %w", u.Email, utils.ErrAlreadyExists)

	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("Error happened during checking user email %v", err)
	}

	// TODO: check create address
	user, err := u.UserDTOToDAL()
	user.ID = 0
	user.IsActive = true
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	if err != nil {
		return nil, err
	}
	newuser, err := dalums.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return dtoums.UserDALToDTO(newuser)
}

// UpdateUser : update exist user.
func UpdateUser(u *dtoums.UserDTO) (*dtoums.UserDTO, error) {
	user, err := u.UserDTOToDAL()
	if err != nil {
		return nil, err
	}
	updateuser, err := dalums.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return dtoums.UserDALToDTO(updateuser)
}

// DeleteUser : delete user by id.
func DeleteUser(id int32) error {
	return dalums.DeleteUser(id)
}
