package domain

import (
	"testing"
)

func Test_market_struct(t *testing.T) {
	dummyMarket := Market{
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
	if dummyMarket.ID != "1" {
		t.Error("Failed while testing the status market struct")
	}
}

func Test_invalid_search_type(t *testing.T) {
	_, err := ValidateSearchType("foo")
	if err == nil {
		t.Error("Failed while validating search type")
	}
}

func Test_no_search_type(t *testing.T) {
	_, err := ValidateSearchType("")
	if err == nil {
		t.Error("Failed while validating search type")
	}
}

func Test_valid_search_type_by_distrito(t *testing.T) {
	_, err := ValidateSearchType("distrito")
	if err != nil {
		t.Error("Failed while validating search type")
	}
}

func Test_valid_search_type_by_regiao5(t *testing.T) {
	_, err := ValidateSearchType("regiao5")
	if err != nil {
		t.Error("Failed while validating search type")
	}
}

func Test_valid_search_type_by_nome_feira(t *testing.T) {
	_, err := ValidateSearchType("nome_feira")
	if err != nil {
		t.Error("Failed while validating search type")
	}
}

func Test_valid_search_type_by_bairro(t *testing.T) {
	_, err := ValidateSearchType("bairro")
	if err != nil {
		t.Error("Failed while validating search type")
	}
}

func Test_invalid_struct(t *testing.T) {
	m := Market{ID: "1"}
	err := ValidateStruct(m)
	if err == nil {
		t.Error("Failed while validating struct")
	}
}

func Test_valid_struct(t *testing.T) {
	m := Market{
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
	err := ValidateStruct(m)
	if err != nil {
		t.Error("Failed while validating struct")
	}
}
