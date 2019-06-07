package main

import(
	"github.com/kas240sx89/samples/golang/profileService/internal/service"
	"github.com/kas240sx89/samples/golang/profileService/internal/db"
)

func main() {

	database := db.NewInMemoryDB()
	var db service.DB
	db = &database
	svc := service.New(db)

	NewServer(svc)
}

