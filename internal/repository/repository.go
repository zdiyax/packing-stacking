package repository

import (
	"packing-stacking/internal/domain"
)

type Repository interface{}

type PacksRepository interface {
	Insert(p domain.Pack) error
	Delete(p domain.Pack) error
	List() ([]domain.Pack, error)
}
