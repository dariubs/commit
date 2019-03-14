package repository

import (
	"fmt"

	"github.com/dariubs/commit/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var err error

func init() {
	cfg, err := config.NewConfig()
	if err != nil {

	}

	conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.HOST,
		cfg.Database.Port,
	)

	DB, err = gorm.Open("postgres", conn)

}

func NewConnection() *gorm.DB {
	return DB
}
