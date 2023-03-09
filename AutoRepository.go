package go_gin

import "context"

//go:generate mockgen -source AutoRepository.go -package mocks -destination mocks/AutoRepository.go
type AutoRepository interface {
	AddAuto(ctx context.Context, auto Auto) error
	GetAllAuto(ctx context.Context) ([]Auto, error)
}
