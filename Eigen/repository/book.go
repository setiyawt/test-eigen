package repository

import (
	"database/sql"
	"myproject/model"
)

type BookRepository interface {
	FetchAll() ([]model.Book, error)
	FetchByID(id int) (*model.Book, error)
	Store(b *model.Book) error
	Update(id int, b *model.Book) error
	Delete(id int) error
}

type bookRepoImpl struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *bookRepoImpl {
	return &bookRepoImpl{db}
}

func (b *bookRepoImpl) FetchAll() ([]model.Book, error) {
	var books []model.Book
	query := "SELECT id, code, title, author, stock FROM books where status ='NotBorrowed'"
	rows, err := b.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Code, &book.Title, &book.Author, &book.Stock)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (b *bookRepoImpl) FetchByID(id int) (*model.Book, error) {
	row := b.db.QueryRow("SELECT id, code, title, author, stock FROM books WHERE id = $1", id)

	var book model.Book
	err := row.Scan(&book.ID, &book.Code, &book.Title, &book.Author, &book.Stock)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *bookRepoImpl) Store(book *model.Book) error {

	_, err := b.db.Exec("INSERT INTO books (code, title, author, stock) VALUES ($1, $2, $3, $4)", book.Code, book.Title, book.Author, book.Stock)
	if err != nil {
		return err
	}
	return nil
}

func (b *bookRepoImpl) Update(id int, book *model.Book) error {
	_, err := b.db.Exec("UPDATE books SET code = $1, title = $2, author = $3, quantity= $4, WHERE id = $5", book.Code, book.Title, book.Author, book.Stock, id)
	if err != nil {
		return err
	}
	return nil
}

func (b *bookRepoImpl) Delete(id int) error {
	_, err := b.db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
