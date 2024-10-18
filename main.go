package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"mi-equipo-backend/src/api/handlers/player"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ginEngine := gin.Default()

	ginEngine.POST("/players", player.CreatePlayer)

	log.Fatalln(ginEngine.Run(":8081"))

}
