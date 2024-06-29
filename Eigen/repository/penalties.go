package repository

import (
	"database/sql"
	"myproject/model"
)

type PenaltiesRepository interface {
	FetchAll() ([]model.Penalties, error)
	FetchByID(id int) (*model.Penalties, error)
	Store(p *model.Penalties) error
	Update(id int, p *model.Penalties) error
	Delete(id int) error
}

type penaltiesRepoImpl struct {
	db *sql.DB
}

func NewPenaltiesRepo(db *sql.DB) *penaltiesRepoImpl {
	return &penaltiesRepoImpl{db}
}

func (p *penaltiesRepoImpl) FetchAll() ([]model.Penalties, error) {
	var penaltiess []model.Penalties
	query := "SELECT id, code_member, penalty_type, penalty_amount, penalty_active, penalty_date, resolved_date FROM penalties where penalty_active = true"
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var penalties model.Penalties
		err := rows.Scan(&penalties.ID, &penalties.CodeMember, &penalties.PenaltyType, &penalties.PenaltyAmount, &penalties.PenaltyActive, &penalties.PenaltyDate, &penalties.ResolveDate)
		if err != nil {
			return nil, err
		}
		penaltiess = append(penaltiess, penalties)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return penaltiess, nil
}

func (p *penaltiesRepoImpl) FetchByID(id int) (*model.Penalties, error) {
	row := p.db.QueryRow("SELECT id, code_member, penalty_type, penalty_amount, penalty_active, penalty_date, resolved_date FROM penalties WHERE id = $1", id)

	var penalties model.Penalties
	err := row.Scan(&penalties.ID, &penalties.CodeMember, &penalties.PenaltyType, &penalties.PenaltyAmount, &penalties.PenaltyActive, &penalties.PenaltyDate, &penalties.ResolveDate)
	if err != nil {
		return nil, err
	}

	return &penalties, nil
}

func (p *penaltiesRepoImpl) Store(penalties *model.Penalties) error {

	_, err := p.db.Exec("INSERT INTO penalties (code_member, penalty_type, penalty_amount, penalty_active, penalty_date, resolved_date) VALUES ($1, $2, $3, $4, $5, $6)", penalties.CodeMember, penalties.PenaltyType, penalties.PenaltyAmount, penalties.PenaltyActive, penalties.PenaltyDate, penalties.ResolveDate)
	if err != nil {
		return err
	}
	return nil
}

func (p *penaltiesRepoImpl) Update(id int, penalties *model.Penalties) error {
	_, err := p.db.Exec("UPDATE penalties SET code_member = $1, penalty_type = $2, penalty_amount = $3, penalty_active = $4, penalty_date = $5, resolved_date = $6 WHERE id = $7", penalties.CodeMember, penalties.PenaltyType, penalties.PenaltyAmount, penalties.PenaltyActive, penalties.PenaltyDate, penalties.ResolveDate, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *penaltiesRepoImpl) Delete(id int) error {
	_, err := p.db.Exec("DELETE FROM penalties WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
