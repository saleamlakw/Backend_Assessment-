package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saleamlakw/LoanTracker/api/route"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	url:=os.Getenv("MONGODB_URL")
	clintOptions:=options.Client().ApplyURI(url)
	client,err:=mongo.Connect(context.TODO(),clintOptions)
	if err!=nil{
		log.Fatal(err)
	}
	err=client.Ping(context.TODO(),nil)
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("connected to mongodb")

	router :=gin.Default()
	route.Route(router,client)
	port:=os.Getenv("PORT")
	router.Run("localhost:"+port)
}