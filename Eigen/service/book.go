package service

import (
	"myproject/model"
	"myproject/repository"
)

type BookService interface {
	FetchAll() ([]model.Book, error)
	FetchByID(id int) (*model.Book, error)
	Store(b *model.Book) error
	Update(id int, b *model.Book) error
	Delete(id int) error
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{bookRepository}
}

func (b *bookService) FetchAll() ([]model.Book, error) {
	books, err := b.bookRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *bookService) FetchByID(id int) (*model.Book, error) {
	book, err := b.bookRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b *bookService) Store(book *model.Book) error {
	err := b.bookRepository.Store(book)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookService) Update(id int, book *model.Book) error {
	err := b.bookRepository.Update(id, book)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookService) Delete(id int) error {
	err := b.bookRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
