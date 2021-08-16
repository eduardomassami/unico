package services

import (
	"errors"
	"testing"
	"unico/internal/core/domain"
	mock_repository "unico/mocks/repository"

	"github.com/golang/mock/gomock"
)

var mockRepo *mock_repository.MockMarketsRepository
var service *Service

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = mock_repository.NewMockMarketsRepository(ctrl)
	service = New(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_error_for_Get(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockRepo.EXPECT().Get("DISTRITO", "1").Return([]domain.Market{}, errors.New("fail to get markets by DISTRITO"))
	// Act
	markets, appError := service.Get("DISTRITO", "1")

	if appError == nil {
		t.Error("Test failed while validating error for get Markets")
	}
	if len(markets) != 0 {
		t.Error("Test failed while validating error for get Markets")
	}
}

func Test_should_return_ok_for_Get(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	dummyMarkets := []domain.Market{
		{ID: "1", Long: "-46550164", Lat: "-23558733", Setcens: "355030885000091", Areap: "3550308005040", Coddist: "87", Distrito: "VILA FORMOSA", CodSubPref: "26", SubPrefe: "ARICANDUVA-FORMOSA-CARRAO", Regiao5: "Leste", Regiao8: "Leste 1", Nome_feira: "VILA FORMOSA", Registro: "4041-0", Logradouro: "RUA MARAGOJIPE", Numero: "S/N", Bairro: "VL FORMOSA", Referencia: "TV RUA PRETORIA"},
		{ID: "1", Long: "-46550164", Lat: "-23558733", Setcens: "355030885000091", Areap: "3550308005040", Coddist: "87", Distrito: "VILA FORMOSA", CodSubPref: "26", SubPrefe: "ARICANDUVA-FORMOSA-CARRAO", Regiao5: "Leste", Regiao8: "Leste 1", Nome_feira: "VILA FORMOSA", Registro: "4041-0", Logradouro: "RUA MARAGOJIPE", Numero: "S/N", Bairro: "VL FORMOSA", Referencia: "TV RUA PRETORIA"},
	}

	mockRepo.EXPECT().Get("DISTRITO", "1").Return(dummyMarkets, nil)
	// Act
	markets, appError := service.Get("DISTRITO", "1")

	if appError != nil {
		t.Error("Test failed while validating error for get Markets")
	}
	if len(markets) != 2 {
		t.Error("Test failed while validating error for get Markets")
	}
}

func Test_should_return_error_for_Post(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	req := domain.Market{
		ID:         "1",
		Long:       "-46550164",
		Lat:        "-23558733",
		Setcens:    "355030885000091",
		Areap:      "3550308005040",
		Coddist:    "87",
		Distrito:   "VILA FORMOSA",
		CodSubPref: "26",
		SubPrefe:   "ARICANDUVA-FORMOSA-CARRAO",
		Regiao5:    "Leste",
		Regiao8:    "Leste 1",
		Nome_feira: "VILA FORMOSA",
		Registro:   "4041-0",
		Logradouro: "RUA MARAGOJIPE",
		Numero:     "S/N",
		Bairro:     "VL FORMOSA",
		Referencia: "TV RUA PRETORIA",
	}
	mockRepo.EXPECT().Save(req).Return(errors.New("fail to create new market"))
	// Act
	appError := service.Post(req)

	if appError == nil {
		t.Error("Test failed while validating error for new Market")
	}
}

func Test_should_return_ok_for_Post(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	req := domain.Market{
		ID:         "1",
		Long:       "-46550164",
		Lat:        "-23558733",
		Setcens:    "355030885000091",
		Areap:      "3550308005040",
		Coddist:    "87",
		Distrito:   "VILA FORMOSA",
		CodSubPref: "26",
		SubPrefe:   "ARICANDUVA-FORMOSA-CARRAO",
		Regiao5:    "Leste",
		Regiao8:    "Leste 1",
		Nome_feira: "VILA FORMOSA",
		Registro:   "4041-0",
		Logradouro: "RUA MARAGOJIPE",
		Numero:     "S/N",
		Bairro:     "VL FORMOSA",
		Referencia: "TV RUA PRETORIA",
	}
	mockRepo.EXPECT().Save(req).Return(nil)
	// Act
	appError := service.Post(req)

	if appError != nil {
		t.Error("Test failed while validating error for new Market")
	}
}

func Test_should_return_error_for_Delete(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockRepo.EXPECT().Remove("1").Return(errors.New("fail to delete market"))
	// Act
	appError := service.Delete("1")

	if appError == nil {
		t.Error("Test failed while validating error for delete Market")
	}
}

func Test_should_return_ok_for_Delete(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockRepo.EXPECT().Remove("1").Return(nil)
	// Act
	appError := service.Delete("1")

	if appError != nil {
		t.Error("Test failed while validating error for delete Market")
	}
}

func Test_should_return_error_for_Put(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	req := domain.Market{
		ID:         "1",
		Long:       "-46550164",
		Lat:        "-23558733",
		Setcens:    "355030885000091",
		Areap:      "3550308005040",
		Coddist:    "87",
		Distrito:   "VILA FORMOSA",
		CodSubPref: "26",
		SubPrefe:   "ARICANDUVA-FORMOSA-CARRAO",
		Regiao5:    "Leste",
		Regiao8:    "Leste 1",
		Nome_feira: "VILA FORMOSA",
		Registro:   "4041-0",
		Logradouro: "RUA MARAGOJIPE",
		Numero:     "S/N",
		Bairro:     "VL FORMOSA",
		Referencia: "TV RUA PRETORIA",
	}
	mockRepo.EXPECT().Alter("1", req).Return(int64(0), errors.New("fail to create new market"))
	// Act
	lines, appError := service.Put("1", req)

	if appError == nil {
		t.Error("Test failed while validating error for new Market")
	}
	if lines != 0 {
		t.Error("Test failed while validating error for new Market")
	}
}

func Test_should_return_ok_for_Put(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	req := domain.Market{
		ID:         "1",
		Long:       "-46550164",
		Lat:        "-23558733",
		Setcens:    "355030885000091",
		Areap:      "3550308005040",
		Coddist:    "87",
		Distrito:   "VILA FORMOSA",
		CodSubPref: "26",
		SubPrefe:   "ARICANDUVA-FORMOSA-CARRAO",
		Regiao5:    "Leste",
		Regiao8:    "Leste 1",
		Nome_feira: "VILA FORMOSA",
		Registro:   "4041-0",
		Logradouro: "RUA MARAGOJIPE",
		Numero:     "S/N",
		Bairro:     "VL FORMOSA",
		Referencia: "TV RUA PRETORIA",
	}
	mockRepo.EXPECT().Alter("1", req).Return(int64(2), nil)
	// Act
	lines, appError := service.Put("1", req)

	if appError != nil {
		t.Error("Test failed while validating error for new Market")
	}
	if lines != 2 {
		t.Error("Test failed while validating error for new Market")
	}
}
