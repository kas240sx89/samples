package main

import (
	db2 "github.com/kas240sx89/samples/profileService/internal/db"
	service2 "github.com/kas240sx89/samples/profileService/internal/service"
)

func main() {

	database := db2.NewInMemoryDB()
	var db service2.DB
	db = &database
	svc := service2.New(db)

	NewServer(svc)
}
