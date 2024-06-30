package repository

import (
	"database/sql"

	"myproject/model"
	"time"
)

type BorrowRepository interface {
	FetchAll() ([]model.Borrowed, error)
	FetchByID(id int) (*model.Borrowed, error)
	Store(b *model.Borrowed) error
	Update(id int, b *model.Borrowed) error
	Delete(id int) error
	GetBorrowedCountByMember(codeMember string) (int, error)
	IsBookCurrentlyBorrowed(codeBook string) (bool, error)
	IsMemberPenalized(codeMember string) (bool, error)
	IsBookExists(codeBook string) (bool, error)
	IsMemberExists(codeMember string) (bool, error)
	isMemberLate(codeMember string) error
	FetchAllReturnBook(codeMember string) ([]model.Borrowed, error)
}

type borrowRepoImpl struct {
	db *sql.DB
}

func NewBorrowRepo(db *sql.DB) *borrowRepoImpl {
	return &borrowRepoImpl{db}
}

func (b *borrowRepoImpl) FetchAll() ([]model.Borrowed, error) {
	var borrowed []model.Borrowed
	query := "SELECT id, code_book, code_member, borroweddate, returneddate, status, late, quantity FROM borrowed"
	rows, err := b.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var borrow model.Borrowed
		err := rows.Scan(
			&borrow.ID,
			&borrow.CodeBook,
			&borrow.CodeMember,
			&borrow.BorrowedDate,
			&borrow.ReturnedDate,
			&borrow.Status,
			&borrow.Late,
			&borrow.Quantity,
		)
		if err != nil {
			return nil, err
		}
		borrowed = append(borrowed, borrow)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return borrowed, nil
}

func (b *borrowRepoImpl) FetchByID(id int) (*model.Borrowed, error) {
	row := b.db.QueryRow("SELECT id, code_book, code_member, borroweddate, returneddate, status, late, quantity FROM borrowed WHERE id = $1", id)

	var borrow model.Borrowed
	err := row.Scan(&borrow.ID, &borrow.CodeBook, &borrow.CodeMember, &borrow.BorrowedDate, &borrow.ReturnedDate, &borrow.Status, &borrow.Late, &borrow.Quantity)
	if err != nil {
		return nil, err
	}

	return &borrow, nil
}

func (b *borrowRepoImpl) IsBookExists(codeBook string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM books
		WHERE code = $1
	`
	var count int
	err := b.db.QueryRow(query, codeBook).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (b *borrowRepoImpl) IsMemberExists(codeMember string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM users
		WHERE code = $1
	`
	var count int
	err := b.db.QueryRow(query, codeMember).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (b *borrowRepoImpl) GetBorrowedCountByMember(codeMember string) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM borrowed
		WHERE code_member = $1 AND status = 'Borrowed'
	`
	var count int
	err := b.db.QueryRow(query, codeMember).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (b *borrowRepoImpl) IsBookCurrentlyBorrowed(codeBook string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM borrowed
		WHERE code_book = $1 AND status = 'Borrowed'
	`
	var count int
	err := b.db.QueryRow(query, codeBook).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (b *borrowRepoImpl) IsMemberPenalized(codeMember string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM penalties
		WHERE code_member = $1 AND penalty_active = TRUE
	`
	var count int
	err := b.db.QueryRow(query, codeMember).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (b *borrowRepoImpl) isMemberLate(codeMember string) error {
	var count int
	query := `
        SELECT COUNT(*)
        FROM borrowed
        WHERE code_member = $1 AND late > 7
    `
	err := b.db.QueryRow(query, codeMember).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	if count == 0 {
		penalty := model.Penalties{
			CodeMember:    codeMember,
			PenaltyType:   "Late Return",
			PenaltyAmount: 50.000,
			PenaltyDate:   time.Now(),
			ResolveDate:   time.Now().AddDate(0, 0, 3),
			PenaltyActive: true,
		}

		_, err = b.db.Exec(`
            INSERT INTO penalties (code_member, penalty_type, penalty_amount, penalty_date, resolve_date, penalty_active)
            VALUES ($1, $2, $3, $4, $5, $6)
        `, penalty.CodeMember, penalty.PenaltyType, penalty.PenaltyAmount, penalty.PenaltyDate, penalty.ResolveDate, penalty.PenaltyActive)

		if err != nil {
			return err
		}
	}

	return nil
}

func (b *borrowRepoImpl) Store(borrow *model.Borrowed) error {
	var lateDays int
	if !borrow.ReturnedDate.IsZero() {
		lateDays = int(time.Now().Sub(borrow.ReturnedDate).Hours() / 24)
	}
	_, err := b.db.Exec("INSERT INTO borrowed (code_book, code_member, borroweddate, returneddate, status, late, quantity) VALUES ($1, $2, $3, $4, $5, $6, $7)", borrow.CodeBook, borrow.CodeMember, borrow.BorrowedDate, borrow.ReturnedDate, borrow.Status, lateDays, borrow.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (b *borrowRepoImpl) Update(id int, borrow *model.Borrowed) error {
	_, err := b.db.Exec("UPDATE borrowed SET code_book = $1, code_member = $2, borroweddate = $3, returneddate = $4, status = $5, late = $6, quantity = $7 WHERE id = $8", borrow.CodeBook, borrow.CodeMember, borrow.BorrowedDate, borrow.ReturnedDate, borrow.Status, borrow.Late, borrow.Quantity, id)
	if err != nil {
		return err
	}
	return nil
}

func (b *borrowRepoImpl) Delete(id int) error {
	_, err := b.db.Exec("DELETE FROM borrowed WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (b *borrowRepoImpl) FetchAllReturnBook(codeMember string) ([]model.Borrowed, error) {
	var borrowed []model.Borrowed

	query := `
        SELECT id, code_book, code_member, borroweddate, returneddate, status, late, quantity 
        FROM borrowed 
        WHERE code_member = $1 AND status = 'Borrowed'
    `

	rows, err := b.db.Query(query, codeMember)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var borrow model.Borrowed
		err := rows.Scan(
			&borrow.ID,
			&borrow.CodeBook,
			&borrow.CodeMember,
			&borrow.BorrowedDate,
			&borrow.ReturnedDate,
			&borrow.Status,
			&borrow.Late,
			&borrow.Quantity,
		)
		if err != nil {
			return nil, err
		}
		borrowed = append(borrowed, borrow)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return borrowed, nil
}
