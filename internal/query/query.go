package query

import (
	"database/sql"
	"fmt"

	"github.com/osag1e/gtc/internal/model"
)

type BookRepo interface {
	InsertBook(book *model.Books) (*model.Books, error)
}

type BookStore struct {
	DB *sql.DB
}

func NewBookStore(db *sql.DB) BookRepo {
	return &BookStore{DB: db}
}

func (ps *BookStore) InsertBook(book *model.Books) (*model.Books, error) {
	query := `
	INSERT INTO store.books (id, title, author, price) 
	VALUES ($1, $2, $3, $4)
	`
	book.ID = model.NewUUID()

	_, err := ps.DB.Exec(query, book.ID, book.Title, book.Author, book.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to insert product: %v", err)
	}
	return book, nil
}
