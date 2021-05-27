package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ogustavobelo/simple-crud-go/controllers"
	"github.com/ogustavobelo/simple-crud-go/core"
	"github.com/ogustavobelo/simple-crud-go/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/acme/autocert"
)

func init() {
	fmt.Printf("on init ")

	services.EnvCheck()
	connectDB()
}

func main() {
	router := gin.Default()
	controllers.InitRoutes(router)

	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}

	server := &http.Server{
		Addr:    ":443",
		Handler: router,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	port := ":" + os.Getenv("SERVER_PORT")
	fmt.Println("Starting server on port: ", port)
	go http.ListenAndServe(port, certManager.HTTPHandler(nil))
	log.Fatal(server.ListenAndServeTLS("", ""))
	// log.Fatal(router.Run(port))
	// log.Fatal(autotls.Run(router, "ogustavobelo.com", "localhost"))
}

// func redirect(w http.ResponseWriter, req *http.Request) {
// 	target := "https://" + req.Host + req.RequestURI

// 	http.Redirect(w, req, target, http.StatusMovedPermanently)
// }

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
