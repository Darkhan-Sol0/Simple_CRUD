package main

import (
	"MyProgy/scr/database"
	"MyProgy/scr/handlers"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	log.Println("Connecting to database")
	db, err := database.ConnectDB(context.Background())
	if err != nil {
		log.Fatalln("DB connection failed: ", err)
	}
	defer db.Close()

	repo := database.NewRepository(db)
	handler := handlers.NewHandler(repo)
	handler.RegHandlers(r)

	log.Println("Starting serveer on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalln("Server failed: ", err)
	}
}
