package service

import (
	"myproject/model"
	"myproject/repository"
)

type PenaltiesService interface {
	FetchAll() ([]model.Penalties, error)
	FetchByID(id int) (*model.Penalties, error)
	Store(p *model.Penalties) error
	Update(id int, p *model.Penalties) error
	Delete(id int) error
}

type penaltiesService struct {
	penaltiesRepository repository.PenaltiesRepository
}

func NewPenaltiesService(penaltiesRepository repository.PenaltiesRepository) PenaltiesService {
	return &penaltiesService{penaltiesRepository}
}

func (p *penaltiesService) FetchAll() ([]model.Penalties, error) {
	penaltiess, err := p.penaltiesRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return penaltiess, nil
}

func (p *penaltiesService) FetchByID(id int) (*model.Penalties, error) {
	penaltiess, err := p.penaltiesRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return penaltiess, nil
}

func (p *penaltiesService) Store(penalties *model.Penalties) error {
	err := p.penaltiesRepository.Store(penalties)
	if err != nil {
		return err
	}

	return nil
}

func (p *penaltiesService) Update(id int, penalties *model.Penalties) error {
	err := p.penaltiesRepository.Update(id, penalties)
	if err != nil {
		return err
	}

	return nil
}

func (p *penaltiesService) Delete(id int) error {
	err := p.penaltiesRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
