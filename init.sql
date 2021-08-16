set global local_infile=true;

CREATE DATABASE IF NOT EXISTS unico;

use unico;

CREATE TABLE IF NOT EXISTS DEINFO_AB_FEIRASLIVRES_2014 (
  ID INT AUTO_INCREMENT PRIMARY KEY,
  `LONG` VARCHAR(255) NOT NULL,
  LAT VARCHAR(255) NOT NULL,
  SETCENS VARCHAR(255) NOT NULL,
  AREAP VARCHAR(255) NOT NULL,
  CODDIST VARCHAR(255) NOT NULL,
  DISTRITO VARCHAR(255) NOT NULL,
  CODSUBPREF VARCHAR(255) NOT NULL,
  SUBPREFE VARCHAR(255) NOT NULL,
  REGIAO5 VARCHAR(255) NOT NULL,
  REGIAO8 VARCHAR(255) NOT NULL,
  NOME_FEIRA VARCHAR(255) NOT NULL,
  REGISTRO VARCHAR(255) NOT NULL,
  LOGRADOURO VARCHAR(255) NOT NULL,
  NUMERO VARCHAR(255),
  BAIRRO VARCHAR(255),
  REFERENCIA VARCHAR(255)
)