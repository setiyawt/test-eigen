package repository

import (
	"database/sql"
	"errors"
	"myproject/model"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
	FetchByID(id int) (*model.User, error)
	FetchAll() ([]model.User, error)
	GetAllMembersWithBorrowedCount() ([]model.MemberBorrow, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	var exists bool
	err := u.db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE code = $1)", user.Code).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("code already exists")
	}

	_, err = u.db.Exec(`INSERT INTO users (code, name, password) VALUES ($1, $2, $3)`, user.Code, user.Name, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) CheckAvail(user model.User) error {
	var count int
	err := u.db.QueryRow("SELECT COUNT(*) FROM users WHERE code = $1 AND password = $2", user.Code, user.Password).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("user not found")
	}
	return nil // TODO: replace this
}

func (u *userRepository) FetchByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT id, code, name, password FROM users WHERE id = $1", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Code, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userRepository) FetchAll() ([]model.User, error) {
	var users []model.User
	query := "SELECT id, code, name FROM users"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Code, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userRepository) GetAllMembersWithBorrowedCount() ([]model.MemberBorrow, error) {
	query := `
		SELECT u.id, u.code, u.name, COUNT(b.id) AS borrow_count
		FROM users u
		LEFT JOIN borrowed b ON u.code = b.code_member AND b.status = 'Borrowed'
		GROUP BY u.id, u.code, u.name
		ORDER BY u.id
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []model.MemberBorrow
	for rows.Next() {
		var member model.MemberBorrow
		err := rows.Scan(&member.ID, &member.Code, &member.Name, &member.BorrowCount)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}
