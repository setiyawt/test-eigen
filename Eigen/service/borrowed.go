package service

import (
	"errors"
	"myproject/model"
	"myproject/repository"
)

type BorrowService interface {
	FetchAll() ([]model.Borrowed, error)
	GetAllMembersWithBorrowedCount() ([]model.User, error)
	FetchByID(id int) (*model.Borrowed, error)
	Store(b *model.Borrowed) error
	Update(id int, b *model.Borrowed) error
	Delete(id int) error
}

type borrowService struct {
	borrowRepository repository.BorrowRepository
}

func NewBorrowService(borrowRepository repository.BorrowRepository) BorrowService {
	return &borrowService{borrowRepository}
}
func (b *borrowService) FetchAll() ([]model.Borrowed, error) {
	borrowed, err := b.borrowRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return borrowed, nil
}

func (b *borrowService) FetchByID(id int) (*model.Borrowed, error) {
	borrow, err := b.borrowRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return borrow, nil

}

func (s *borrowService) GetAllMembersWithBorrowedCount() ([]model.User, error) {
	borrowed, err := s.borrowRepository.GetAllMembersWithBorrowedCount()
	if err != nil {
		return nil, err
	}

	return borrowed, nil
}

func (b *borrowService) Store(borrow *model.Borrowed) error {

	bookExists, err := b.borrowRepository.IsBookExists(borrow.CodeBook)
	if err != nil {
		return err
	}
	if !bookExists {
		return errors.New("book does not exist")
	}

	memberExists, err := b.borrowRepository.IsMemberExists(borrow.CodeMember)
	if err != nil {
		return err
	}
	if !memberExists {
		return errors.New("member does not exist")
	}
	isPenalized, err := b.borrowRepository.IsMemberPenalized(borrow.CodeMember)
	if err != nil {
		return err
	}
	if isPenalized {
		return errors.New("member is currently penalized and cannot borrow books")
	}

	// Check if the member is already borrowing 2 books
	borrowCount, err := b.borrowRepository.GetBorrowedCountByMember(borrow.CodeMember)
	if err != nil {
		return err
	}
	if borrowCount >= 2 {
		return errors.New("member cannot borrow more than 2 books at the same time")
	}

	// Check if the book is currently borrowed
	isBookBorrowed, err := b.borrowRepository.IsBookCurrentlyBorrowed(borrow.CodeBook)
	if err != nil {
		return err
	}
	if isBookBorrowed {
		return errors.New("the book is already borrowed by another member")
	}

	// Save the borrowing record
	err = b.borrowRepository.Store(borrow)
	if err != nil {
		return err
	}

	return nil
}

func (b *borrowService) Update(id int, borrow *model.Borrowed) error {
	err := b.borrowRepository.Update(id, borrow)
	if err != nil {
		return err
	}

	return nil
}

func (b *borrowService) Delete(id int) error {
	err := b.borrowRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
