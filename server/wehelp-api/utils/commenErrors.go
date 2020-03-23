package utils

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// ErrRecordNotFound :
var ErrRecordNotFound = gorm.ErrRecordNotFound

// ErrAlreadyExists :
var ErrAlreadyExists = errors.New("already exists")
