package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

const (
	dsnKey = "DSN"
)

var DB *gorm.DB

func SetupDB() error {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv(dsnKey)), &gorm.Config{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&Article{}, &ArticleLine{}, &ArticleWord{})
	if err != nil {
		return err
	}

	res := DB.First(&ArticleWord{})
	// if there's already data in the db, do not populate
	if res.Error == nil {
		return nil
	}

	err = populateDB()
	if err != nil {
		return err
	}

	return nil
}
