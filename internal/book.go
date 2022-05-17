package internal

import (
	"context"
	"time"
)

type BookId int64

type Book struct {
	Bid         BookId    `json:"bid"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	FinishTime  time.Time `json:"finishTime"`
}

type BookRepository interface {
	CreateBook(ctx context.Context, book *Book) (BookId, error)
	RetrieveBook(ctx context.Context, bid BookId) (*Book, error)
	UpdateBook(ctx context.Context, book *Book) error
	DeleteBook(ctx context.Context, bid BookId) error
	ListBook(ctx context.Context, offset int64, limit int64) ([]*Book, error)
}
