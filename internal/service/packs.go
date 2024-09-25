package service

import (
	"packing-stacking/internal/domain"
	"packing-stacking/internal/repository"
	_map "packing-stacking/internal/repository/map"
)

type PacksService interface {
	Add(p domain.Pack) error
	Remove(p domain.Pack) error
	List() ([]domain.Pack, error)
	Calculate(quantity int64) (map[domain.Pack]int64, error)
}

type packsService struct {
	repo repository.PacksRepository
}

func (svc packsService) Add(p domain.Pack) error {
	return svc.repo.Insert(p)
}

func (svc packsService) Remove(p domain.Pack) error {
	return svc.repo.Delete(p)
}

func (svc packsService) List() ([]domain.Pack, error) {
	return svc.repo.List()
}

func (svc packsService) Calculate(quantity int64) (map[domain.Pack]int64, error) {

	panic("implement me")
}

func NewPacksService() PacksService {
	return &packsService{repo: _map.NewPacksMapRepository()}
}
