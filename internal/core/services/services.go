package services

import (
	"unico/internal/core/domain"
	"unico/internal/core/ports"
)

type Service struct {
	marketsRepository ports.MarketsRepository
}

func New(marketsRepository ports.MarketsRepository) *Service {
	return &Service{
		marketsRepository: marketsRepository,
	}
}

func (srv *Service) Get(search string, id string) ([]domain.Market, error) {
	markets, err := srv.marketsRepository.Get(search, id)
	if err != nil {
		return []domain.Market{}, err
	}

	return markets, nil
}

func (srv *Service) Post(m domain.Market) error {
	if err := srv.marketsRepository.Save(m); err != nil {
		return err
	}

	return nil
}

func (srv *Service) Delete(id string) error {
	if err := srv.marketsRepository.Remove(id); err != nil {
		return err
	}

	return nil
}

func (srv *Service) Put(id string, m domain.Market) (int64, error) {
	market, err := srv.marketsRepository.Alter(id, m)
	if err != nil {
		return 0, err
	}

	return market, nil
}
