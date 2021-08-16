package repository

import (
	"database/sql"
	"errors"
	"unico/internal/core/domain"
	"unico/logger"
)

type MarketRepositoryDb struct {
	client *sql.DB
}

func (d MarketRepositoryDb) Get(search string, id string) ([]domain.Market, error) {
	marketSql := "SELECT * from DEINFO_AB_FEIRASLIVRES_2014 WHERE " + search + " = ?"
	var m domain.Market

	res, err := d.client.Query(marketSql, id)
	if err != nil {
		logger.ErrorLog.Println("fail to get markets by DISTRITO")
		return []domain.Market{}, errors.New("fail to get markets by DISTRITO")
	}

	var markets []domain.Market

	for res.Next() {

		err := res.Scan(
			&m.ID,
			&m.Areap,
			&m.Bairro,
			&m.Coddist,
			&m.CodSubPref,
			&m.Distrito,
			&m.Lat,
			&m.Logradouro,
			&m.Long,
			&m.Nome_feira,
			&m.Numero,
			&m.Referencia,
			&m.Regiao5,
			&m.Regiao8,
			&m.Registro,
			&m.Setcens,
			&m.SubPrefe)

		if err != nil {
			logger.ErrorLog.Println("fail to get markets by DISTRITO")
			return []domain.Market{}, errors.New("fail to get markets by DISTRITO")
		}
		markets = append(markets, m)
	}

	if len(markets) == 0 {
		logger.SuccessLog.Println("no market found")
		return []domain.Market{}, errors.New("no market found")
	}

	return markets, nil
}

func (d MarketRepositoryDb) Save(m domain.Market) error {
	marketSql := `INSERT INTO DEINFO_AB_FEIRASLIVRES_2014 VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := d.client.Exec(marketSql,
		m.ID,
		m.Long,
		m.Lat,
		m.Setcens,
		m.Areap,
		m.Coddist,
		m.Distrito,
		m.CodSubPref,
		m.SubPrefe,
		m.Regiao5,
		m.Regiao8,
		m.Nome_feira,
		m.Registro,
		m.Logradouro,
		m.Numero,
		m.Bairro,
		m.Referencia,
	)

	if err != nil {
		logger.ErrorLog.Println("fail to create new market")
		return errors.New("fail to create new market")
	}

	return nil
}

func (d MarketRepositoryDb) Remove(id string) error {
	marketSql := `DELETE FROM DEINFO_AB_FEIRASLIVRES_2014 WHERE ID = ?`

	_, err := d.client.Query(marketSql, id)

	if err != nil {
		logger.ErrorLog.Println("fail to delete market")
		return errors.New("fail to delete market")
	}

	return nil
}

func (d MarketRepositoryDb) Alter(id string, m domain.Market) (int64, error) {
	marketSql := "UPDATE DEINFO_AB_FEIRASLIVRES_2014 SET `LONG` = ?, LAT = ?, SETCENS = ?, AREAP = ?, CODDIST = ?, DISTRITO = ?, CODSUBPREF = ?, SUBPREFE = ?, REGIAO5 = ?, REGIAO8 = ?, NOME_FEIRA = ?, REGISTRO = ?, LOGRADOURO = ?, NUMERO = ?, BAIRRO = ?, REFERENCIA = ? WHERE ID = ?"

	result, err := d.client.Exec(marketSql,
		m.Long,
		m.Lat,
		m.Setcens,
		m.Areap,
		m.Coddist,
		m.Distrito,
		m.CodSubPref,
		m.SubPrefe,
		m.Regiao5,
		m.Regiao8,
		m.Nome_feira,
		m.Registro,
		m.Logradouro,
		m.Numero,
		m.Bairro,
		m.Referencia,
		id,
	)

	if err != nil {
		logger.ErrorLog.Println("fail to update market")
		return 0, errors.New("fail to update market")
	}

	rows, err := result.RowsAffected()

	if err != nil {
		logger.ErrorLog.Println("fail to update market")
		return 0, errors.New("fail to update market")
	}

	return rows, nil
}

func NewMarketRepositoryDb(dbClient *sql.DB) MarketRepositoryDb {
	return MarketRepositoryDb{dbClient}
}
