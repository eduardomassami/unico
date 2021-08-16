package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unico/internal/core/domain"
	ports "unico/mocks/service"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var router *mux.Router
var ch HTTPHandler
var mockMarketsService *ports.MockMarketsService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockMarketsService = ports.NewMockMarketsService(ctrl)
	ch = HTTPHandler{mockMarketsService}
	router = mux.NewRouter()
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_markets_with_status_200_for_get_markets(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	dummyMarkets := []domain.Market{
		{ID: "1", Long: "-46550164", Lat: "-23558733", Setcens: "355030885000091", Areap: "3550308005040", Coddist: "87", Distrito: "VILA FORMOSA", CodSubPref: "26", SubPrefe: "ARICANDUVA-FORMOSA-CARRAO", Regiao5: "Leste", Regiao8: "Leste 1", Nome_feira: "VILA FORMOSA", Registro: "4041-0", Logradouro: "RUA MARAGOJIPE", Numero: "S/N", Bairro: "VL FORMOSA", Referencia: "TV RUA PRETORIA"},
		{ID: "1", Long: "-46550164", Lat: "-23558733", Setcens: "355030885000091", Areap: "3550308005040", Coddist: "87", Distrito: "VILA FORMOSA", CodSubPref: "26", SubPrefe: "ARICANDUVA-FORMOSA-CARRAO", Regiao5: "Leste", Regiao8: "Leste 1", Nome_feira: "VILA FORMOSA", Registro: "4041-0", Logradouro: "RUA MARAGOJIPE", Numero: "S/N", Bairro: "VL FORMOSA", Referencia: "TV RUA PRETORIA"},
	}
	router.HandleFunc("/market/{searchType}/{id}", ch.Get)
	mockMarketsService.EXPECT().Get("DISTRITO", "1").Return(dummyMarkets, nil)
	request, _ := http.NewRequest(http.MethodGet, "/market/distrito/1", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message_for_get_markets_when_invalid_search_type(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market/{searchType}/{id}", ch.Get)

	request, _ := http.NewRequest(http.MethodGet, "/market/foo/1", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message_for_get_markets(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market/{searchType}/{id}", ch.Get)
	mockMarketsService.EXPECT().Get("DISTRITO", "1").Return(nil, errors.New("fail to get markets by DISTRITO"))
	request, _ := http.NewRequest(http.MethodGet, "/market/distrito/1", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_markets_with_status_200_for_post_markets(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market", ch.Post)
	dummyMarket := domain.Market{
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

	jsonValue, _ := json.Marshal(dummyMarket)
	mockMarketsService.EXPECT().Post(dummyMarket).Return(nil)
	request, _ := http.NewRequest(http.MethodPost, "/market", bytes.NewBuffer(jsonValue))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusCreated {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_markets_with_status_400_for_post_markets_when_market_struct_invalid(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market", ch.Post)
	dummyMarket := domain.Market{
		ID:   "1",
		Long: "-46550164",
	}
	jsonValue, _ := json.Marshal(dummyMarket)
	request, _ := http.NewRequest(http.MethodPost, "/market", bytes.NewBuffer(jsonValue))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message_for_post_markets(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market", ch.Post)
	dummyMarket := domain.Market{
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
	jsonValue, _ := json.Marshal(dummyMarket)
	mockMarketsService.EXPECT().Post(dummyMarket).Return(errors.New("fail to create new market"))
	request, _ := http.NewRequest(http.MethodPost, "/market", bytes.NewBuffer(jsonValue))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_400_with_error_message_for_post_markets(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market", ch.Post)
	request, _ := http.NewRequest(http.MethodPost, "/market", strings.NewReader("trouble will find me"))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusBadRequest {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_markets_with_status_200_for_delete_markets(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market/{id}", ch.Delete)
	mockMarketsService.EXPECT().Delete("1").Return(nil)
	request, _ := http.NewRequest(http.MethodDelete, "/market/1", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message_for_delete_markets(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market/{id}", ch.Delete)
	mockMarketsService.EXPECT().Delete("1").Return(errors.New("fail to delete market"))
	request, _ := http.NewRequest(http.MethodDelete, "/market/1", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_markets_with_status_200_for_put_markets(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market/{id}", ch.Put)
	dummyMarket := domain.Market{
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

	jsonValue, _ := json.Marshal(dummyMarket)
	mockMarketsService.EXPECT().Put("1", dummyMarket).Return(int64(1), nil)
	request, _ := http.NewRequest(http.MethodPut, "/market/1", bytes.NewBuffer(jsonValue))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message_for_put_markets_when_market_struct_invalid(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market/{id}", ch.Put)
	dummyMarket := domain.Market{
		ID:   "1",
		Long: "-46550164",
	}
	jsonValue, _ := json.Marshal(dummyMarket)

	request, _ := http.NewRequest(http.MethodPut, "/market/1", bytes.NewBuffer(jsonValue))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message_for_put_markets(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market/{id}", ch.Put)
	dummyMarket := domain.Market{
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
	jsonValue, _ := json.Marshal(dummyMarket)
	mockMarketsService.EXPECT().Put("1", dummyMarket).Return(int64(0), errors.New("fail to create new market"))
	request, _ := http.NewRequest(http.MethodPut, "/market/1", bytes.NewBuffer(jsonValue))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_400_with_error_message_for_put_markets(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	router.HandleFunc("/market/{id}", ch.Put)
	request, _ := http.NewRequest(http.MethodPut, "/market/1", strings.NewReader("trouble will find me"))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusBadRequest {
		t.Error("Failed while testing the status code")
	}
}
