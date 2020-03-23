package ums

import (
	"database/sql"
	"wehelp-api/db"
)

// UserDAL : data access layer  (user) table.
type UserDAL struct {
	ID          int32          `json:"id" gorm:"column:id;primary_key:true"`
	Email       string         `json:"email" gorm:"column:email"`
	Name        string         `json:"name" gorm:"column:name"`
	PasswordMd5 string         `json:"password_md5" gorm:"column:password_md5"`
	Phone       sql.NullString `json:"phone" gorm:"column:phone"`
	Bio         sql.NullString `json:"bio" gorm:"column:bio"`
	Address     sql.NullInt32  `json:"address" gorm:"column:address"`
	IsActive    bool           `json:"is_active" gorm:"column:is_active"`
	CreatedAt   int64          `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   int64          `json:"updated_at" gorm:"column:updated_at"`
}

// TableName sets the insert table name for this struct type
func (u *UserDAL) TableName() string {
	return "user"
}

// GetAllUseres : get all useres.
func GetAllUseres() []*UserDAL {
	useres := []*UserDAL{}
	db.DB().Find(&useres)
	return useres
}

// GetUser : get one user by id.
func GetUser(id int32) (*UserDAL, error) {
	u := &UserDAL{}
	result := db.DB().First(u, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

// GetUserByEmail : get one user by email.
func GetUserByEmail(email string) (*UserDAL, error) {
	u := &UserDAL{}
	result := db.DB().First(u, UserDAL{Email: email})
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

// CreateUser : create new user.
func CreateUser(u *UserDAL) (*UserDAL, error) {
	result := db.DB().Create(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

// UpdateUser : update exist user.
func UpdateUser(u *UserDAL) (*UserDAL, error) {
	_, err := GetUser(u.ID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

// DeleteUser : delete user by id.
func DeleteUser(id int32) error {
	u, err := GetUser(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(u)
	return result.Error
}
