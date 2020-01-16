package models

import (
	srv "intab-core/services"
	"log"
)

//Migrate 迁移数据库
func Migrate() {
	srv.DB().LogMode(true)
	//Add New Model Name Here
	if err := srv.DB().AutoMigrate(
		&User{},
		&UserDetail{},
		&History{},
		&Snapshot{},
		&Document{},
		&Resource{},
		//&Sharekey{},
	).Error; err != nil {
		log.Println("failed to migrate database:")
		panic(err)
	}
}
