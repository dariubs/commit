package model

import (
	"github.com/dariubs/commit/repository"
	"github.com/jinzhu/gorm"
)

// DB variable
var DB *gorm.DB

func init() {
	DB = repository.NewConnection()
}
