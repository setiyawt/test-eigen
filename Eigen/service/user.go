package service

import (
	"myproject/model"
	"myproject/repository"
)

type UserService interface {
	Login(user model.User) error
	Register(user model.User) error
	FetchAll() ([]model.User, error)
	CheckPassLength(pass string) bool
	CheckPassAlphabet(pass string) bool
	GetAllMembersWithBorrowedCount() ([]model.MemberBorrow, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Login(user model.User) error {
	err := s.userRepository.CheckAvail(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Register(user model.User) error {
	err := s.userRepository.Add(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) CheckPassLength(pass string) bool {
	if len(pass) <= 5 {
		return true
	}

	return false
}

func (s *userService) CheckPassAlphabet(pass string) bool {
	for _, charVariable := range pass {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {
			return false
		}
	}
	return true
}

func (s *userService) FetchAll() ([]model.User, error) {
	users, err := s.userRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) GetAllMembersWithBorrowedCount() ([]model.MemberBorrow, error) {
	borrowed, err := s.userRepository.GetAllMembersWithBorrowedCount()
	if err != nil {
		return nil, err
	}

	return borrowed, nil
}
