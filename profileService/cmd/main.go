package main

import (
	//"github.com/kas240sx89/samples/profileService/config"
	"github.com/kas240sx89/samples/profileService/internal/db"
	"github.com/kas240sx89/samples/profileService/internal/service"
)

func main() {

	database := db.NewInMemoryDB()
	//var sDB service.DB
	//sDB = &database
	svc := service.New(&database)

	NewServer(svc)
}

//func GetDB(cfg *config.Config) service.DB {
//	var sDB service.DB
//	switch cfg.ProfileDBType {
//	case "inmemory":
//	}
//
//}
