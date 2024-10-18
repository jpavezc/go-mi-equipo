package player

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mi-equipo-backend/src/api/internal/domain"
	"os"
	"time"
)

func CreatePlayerService(player domain.Player) (id interface{}, err error) {

	player.CreatedAt = time.Now().UTC()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalln("Error connecting to Mongo", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	collection := client.Database("go-l").Collection("players")
	insertedResult, err := collection.InsertOne(ctx, player)
	if err != nil {
		log.Fatalln("Error inserting into the database", err)

	}
	return insertedResult.InsertedID, nil
}
