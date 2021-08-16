package repository

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func Test_get_shoud_return_markets(t *testing.T) {
	db, mock := NewMock()
	repo := NewMarketRepositoryDb(db)
	defer func() {
		repo.client.Close()
	}()
	search := "DISTRITO"
	marketSql := "SELECT * from DEINFO_AB_FEIRASLIVRES_2014 WHERE " + search + " = ?"
	rows := sqlmock.NewRows([]string{"ID", "Long", "Lat", "Setcens", "Areap", "Coddist", "Distrito", "CodSubPref", "SubPrefe", "Regiao5", "Regiao8", "Nome_feira", "Registro", "Logradouro", "Numero", "Bairro", "Referencia"}).
		AddRow("1", "-46550164", "-23558733", "355030885000091", "3550308005040", "87", "VILA FORMOSA", "26", "ARICANDUVA-FORMOSA-CARRAO", "Leste", "Leste 1", "VILA FORMOSA", "4041-0", "RUA MARAGOJIPE", "S/N", "VL FORMOSA", "TV RUA PRETORIA")

	mock.ExpectQuery(regexp.QuoteMeta(marketSql)).WithArgs("1").WillReturnRows(rows)

	markets, err := repo.Get(search, "1")
	if markets == nil {
		t.Error("Error while testing Get")
	}
	if err != nil {
		t.Error("Error while testing Get")
	}
}
