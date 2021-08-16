package domain

import (
	"errors"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type Market struct {
	ID         string `json:"id" validate:"required"`
	Long       string `json:"long" validate:"required"`
	Lat        string `json:"lat" validate:"required"`
	Setcens    string `json:"setcens" validate:"required"`
	Areap      string `json:"areap" validate:"required"`
	Coddist    string `json:"coddist" validate:"required"`
	Distrito   string `json:"distrito" validate:"required"`
	CodSubPref string `json:"codsubpref" validate:"required"`
	SubPrefe   string `json:"subprefe" validate:"required"`
	Regiao5    string `json:"regiao5" validate:"required"`
	Regiao8    string `json:"regiao8" validate:"required"`
	Nome_feira string `json:"nome_feira" validate:"required"`
	Registro   string `json:"registro" validate:"required"`
	Logradouro string `json:"logradouro" validate:"required"`
	Numero     string `json:"numero"`
	Bairro     string `json:"bairro"`
	Referencia string `json:"referencia"`
}

func ValidateSearchType(searchType string) (string, error) {
	search := strings.ToUpper(searchType)
	if search == "" {
		return "", errors.New("invalid search type")
	}

	switch search {
	case "DISTRITO":
		return strings.ToUpper(searchType), nil
	case "REGIAO5":
		return strings.ToUpper(searchType), nil
	case "NOME_FEIRA":
		return strings.ToUpper(searchType), nil
	case "BAIRRO":
		return strings.ToUpper(searchType), nil
	default:
		return "", errors.New("invalid search type")
	}
}

func ValidateStruct(m Market) error {
	validate = validator.New()
	err := validate.Struct(m)
	if err != nil {
		return err
	}
	return nil
}
