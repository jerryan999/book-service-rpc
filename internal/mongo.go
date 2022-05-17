package internal

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	bookCollection = "books"
)

type MongoBookRepository struct {
	counter    BookId // increment book id
	mu         sync.Mutex
	collection *mongo.Collection
}

func NewMongoBookRepository(db *mongo.Database) *MongoBookRepository {
	return &MongoBookRepository{
		collection: db.Collection(bookCollection),
	}
}

func (r *MongoBookRepository) CreateBook(ctx context.Context, book *Book) (BookId, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.counter += 1

	book.Bid = r.counter
	_, err := r.collection.InsertOne(ctx, book)
	if err != nil {
		return 0, err
	}
	return r.counter, err
}

func (r *MongoBookRepository) RetrieveBook(ctx context.Context, bid BookId) (*Book, error) {
	var book Book
	err := r.collection.FindOne(ctx, map[string]BookId{"bid": bid}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *MongoBookRepository) UpdateBook(ctx context.Context, book *Book) error {
	_, err := r.collection.UpdateOne(ctx, map[string]BookId{"bid": BookId(book.Bid)}, map[string]interface{}{"$set": book})
	return err
}

func (r *MongoBookRepository) DeleteBook(ctx context.Context, bid BookId) error {
	_, err := r.collection.DeleteOne(ctx, map[string]BookId{"bid": bid})
	return err
}

func (r *MongoBookRepository) ListBook(ctx context.Context, offset int64, limit int64) ([]*Book, error) {
	var books []*Book
	cursor, err := r.collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}
