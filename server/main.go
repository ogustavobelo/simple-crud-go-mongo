package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ogustavobelo/simple-crud-go/controllers"
	"github.com/ogustavobelo/simple-crud-go/core"
	"github.com/ogustavobelo/simple-crud-go/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	services.EnvCheck()
	connectDB()
}

func main() {
	router := gin.Default()
	controllers.InitRoutes(router)

	port := ":" + os.Getenv("SERVER_PORT")
	fmt.Printf("Starting server on port %v...", port)
	log.Fatal(router.Run(port))
}

func connectDB() {
	databasePort := os.Getenv("DATABASE_PORT")
	fmt.Printf("connecting database on Port %v ...", databasePort)
	mongoURI := fmt.Sprintf("mongodb://mongodb:%v/", databasePort)
	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection(core.USERS)
	controllers.SetCollection(collection)

}
