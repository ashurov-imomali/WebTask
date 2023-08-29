package db

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/pkg/models"
	"os"
)

func ConnectionToDb(path string) (*gorm.DB, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var Db models.DbSettings
	err = json.Unmarshal(bytes, &Db)
	if err != nil {
		return nil, err
	}
	dsn := "host=" + Db.Host + " port=" + Db.Port + " user=" + Db.UserName + " password=" + Db.Password + " dbname=" + Db.Database
	//dsn := "postgres://" + Db.UserName + ":" + Db.Password + "@" + Db.Host + ":" + Db.Port + "/" + Db.Database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
