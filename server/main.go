package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ogustavobelo/simple-crud-go/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	connectDB()
}

func main() {
	router := gin.Default()
	controllers.InitRoutes(router)
	log.Fatal(router.Run(":3000"))
}

func connectDB() {
	fmt.Println("connecting...")
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017/")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("chat").Collection("users2")
	controllers.SetCollection(collection)

}
