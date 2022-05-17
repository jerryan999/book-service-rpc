package main

import (
	"context"
	"log"
	"net"
	"time"

	internal "github.com/jerryan999/book-service/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.Println("Starting listening on port 8080")
	port := ":8080"

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Mongo Repository
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("testing")

	// create a new mongo repository
	var repository internal.BookRepository = internal.NewMongoBookRepository(db)
	srv := internal.NewRPCServer(repository)

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
