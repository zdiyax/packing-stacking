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
	Calculate(quantity int) (map[int]int, error)
}

type packsService struct {
	repo repository.PacksRepository
}

func NewPacksService() PacksService {
	return &packsService{repo: _map.NewPacksMapRepository()}
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

func (svc packsService) Calculate(quantity int) (map[int]int, error) {
	packs, err := svc.repo.List()
	if err != nil {
		return nil, err
	}

	calculation := make(map[int]int)
	minPacks := make(map[int]int)       // Map to track the minimum combination
	minPacksCount := int(^uint(0) >> 1) // Set to a large number (infinity)
	minExcess := int(^uint(0) >> 1)     // Set to a large number to track the smallest overfill

	var backtrack func(remainingItems int, packIdx int, currentCombo map[int]int, packCount int)

	backtrack = func(remainingItems int, packIdx int, currentCombo map[int]int, packCount int) {
		// If remaining items are <= 0 (we've packed enough or exceeded), check if this is a better solution
		if remainingItems <= 0 {
			excess := -remainingItems // How much we've overpacked
			// Prefer the combination that minimizes the number of packs, and if equal, the smallest overfill
			if packCount < minPacksCount || (packCount == minPacksCount && excess < minExcess) {
				minPacksCount = packCount
				minExcess = excess
				// Copy the current combination to the result
				for k, v := range currentCombo {
					minPacks[k] = v
				}
			}
			return
		}

		// Try all packs starting from the current index
		for i := packIdx; i < len(packs); i++ {
			pack := packs[i]
			if remainingItems >= pack.Quantity {
				currentCombo[pack.Quantity]++
				backtrack(remainingItems-pack.Quantity, i, currentCombo, packCount+1)
				currentCombo[pack.Quantity]-- // Backtrack
				if currentCombo[pack.Quantity] == 0 {
					delete(currentCombo, pack.Quantity) // Clean up map
				}
			}
		}
	}

	// Start backtracking
	backtrack(quantity, 0, calculation, 0)

	return minPacks, nil
}
