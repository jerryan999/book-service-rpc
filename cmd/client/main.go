package main

import (
	"context"
	"fmt"
	"log"

	api "github.com/jerryan999/book-service/api/v1"
	gRPC "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var serverAddress = "localhost:8080"

func main() {
	conn, err := gRPC.Dial(serverAddress, gRPC.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := api.NewBookServiceClient(conn)

	// add book
	bookDTO := &api.Book{
		Title:       "Go Programming",
		Author:      "John Doe",
		Description: "Go is a programming language",
		Language:    "English",
		FinishTime:  timestamppb.Now(),
	}
	resCreate, err := client.CreateBook(context.Background(), &api.CreateBookRequest{Book: bookDTO})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
	}
	log.Printf("Book created with bid: %d\n", resCreate.Bid)

	// retrieve a book
	var bid int64 = 1
	resRetrieve, err := client.RetrieveBook(context.Background(), &api.RetrieveBookRequest{Bid: bid})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
	} else {
		log.Printf("Book retrieved: %v\n", resRetrieve.Book.String())
	}

	// update a book
	var bidUpdate int64 = 3
	bookUpdate := &api.Book{
		Bid:         bidUpdate,
		Title:       "Go Programming-updated",
		Author:      "John Doe",
		Description: "Go is a programming language",
		Language:    "English",
		FinishTime:  timestamppb.Now(),
	}
	_, err = client.UpdateBook(context.Background(), &api.UpdateBookRequest{Book: bookUpdate})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
		fmt.Println(errStatus.Code())
	} else {
		log.Printf("Book updated: %v\n", bookUpdate.String())
	}

	// delete book
	var bidDelete int64 = 2
	_, err = client.DeleteBook(context.Background(), &api.DeleteBookRequest{Bid: bidDelete})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
		fmt.Println(errStatus.Code())
	} else {
		log.Printf("Book deleted bid: %v\n", bidDelete)
	}

	// list book
	resList, err := client.ListBook(context.Background(), &api.ListBookRequest{})
	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
		fmt.Println(errStatus.Code())
	} else {
		log.Printf("Book list: %v\n", resList.Books)
	}

}
