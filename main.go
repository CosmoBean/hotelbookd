package main

import (
	"github.com/CosmoBean/hotelbookd/db"
	"github.com/CosmoBean/hotelbookd/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Panic("error Loading the env file: ", err)
	}

	server.Init()       //default port :8080
	db.GetMongoClient() //Init DBInstance

}
