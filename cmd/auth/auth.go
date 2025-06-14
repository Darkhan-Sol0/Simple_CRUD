package main

import (
	"MyProgy/infrastructure/database"
	"MyProgy/internal/datasource"
	"MyProgy/internal/web"
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

	repo := datasource.NewRepository(db)
	handler := web.NewHandler(repo)
	handler.RegHandlers(r)

	log.Println("Starting serveer on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalln("Server failed: ", err)
	}
}
