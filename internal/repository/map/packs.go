package _map

import (
	"packing-stacking/internal/repository"
	"sort"
	"sync"

	"packing-stacking/internal/domain"
)

type packsMapRepository struct {
	m *sync.Map
}

func NewPacksMapRepository() repository.PacksRepository {
	m := &sync.Map{}
	m.Store(int(100), &domain.Pack{Quantity: 100})
	m.Store(int(250), &domain.Pack{Quantity: 250})
	m.Store(int(500), &domain.Pack{Quantity: 500})

	return &packsMapRepository{m: m}
}

func (r *packsMapRepository) Insert(p domain.Pack) error {
	r.m.Store(p.Quantity, &p)

	return nil
}

func (r *packsMapRepository) Delete(p domain.Pack) error {
	r.m.Delete(p.Quantity)

	return nil
}

func (r *packsMapRepository) List() ([]domain.Pack, error) {
	var result []domain.Pack

	r.m.Range(func(key, value any) bool {
		pack := value.(*domain.Pack)

		result = append(result, *pack)
		return true
	})

	sort.Slice(result, func(i, j int) bool {
		return result[i].Quantity > result[j].Quantity
	})

	return result, nil
}
