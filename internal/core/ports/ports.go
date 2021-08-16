package ports

import "unico/internal/core/domain"

//go:generate mockgen -destination=../../../mocks/repository/mockMarketsRepository.go -package=ports unico/internal/core/ports MarketsRepository
type MarketsRepository interface {
	Get(search string, id string) ([]domain.Market, error)
	Save(domain.Market) error
	Remove(id string) error
	Alter(id string, m domain.Market) (int64, error)
}

//go:generate mockgen -destination=../../../mocks/service/mockMarketsService.go -package=ports unico/internal/core/ports MarketsService
type MarketsService interface {
	Get(search string, id string) ([]domain.Market, error)
	Post(domain.Market) error
	Delete(id string) error
	Put(id string, m domain.Market) (int64, error)
}
