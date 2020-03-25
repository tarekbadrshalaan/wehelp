package ums

import "wehelp-api/db"

// GetUserByEmailPassword : get one user by email and password.
func GetUserByEmailPassword(email, passmd5 string) (*UserDAL, error) {
	u := &UserDAL{}
	result := db.DB().First(u, UserDAL{Email: email, PasswordMd5: passmd5})
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
