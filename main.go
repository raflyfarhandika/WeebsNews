package main

import "weebsnews/database"

func main() {

	DB := database.DBStart()
	database.DBMigrate()
	defer DB.Close()

}
