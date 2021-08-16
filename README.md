# CRUD feito com GOLANG e MYSQL para controle de feiras livres da cidade de São Paulo
Primeiro passo é criar um arquivo .env para definir as variáveis de ambiente do projeto:
```
SERVER_ADDRESS=localhost
SERVER_PORT=8000
DB_USER=admin
DB_PASSWD=123456
DB_ADDR=localhost
DB_PORT=3306
DB_NAME=unico
LOG_FILE_LOCATION=../../logs
APP_NAME=unico
FILE_NAME=DEINFO_AB_FEIRASLIVRES_2014
```
Segundo passo editar o start.sh e colocar as variáveis para importação da planilha
```
DB_USER=admin \
DB_PASSWD=123456 \
DB_NAME=unico \
FILE_NAME=DEINFO_AB_FEIRASLIVRES_2014 \
```
Com todas as variáveis criadas, execute o start.sh
```
$ ./start.sh
```
O App irá rodar no SERVER_ADDRESS:SERVER_PORT definidos no primeiro passo

# curl's de exemplo
Busca de feira
```
$ curl --location --request GET 'localhost:8000/market/{searchType}/{id}'
```
OBS.: Busca de feiras pode ser feita por distrito, regiao5, nome_feira, bairro, para tal coloque o critério em searchType e o que deseja buscar, ex:
```
$ curl --location --request GET 'localhost:8000/market/distrito/SACOMA'
```

Cadastro de nova feira
```
$ curl --location --request POST 'localhost:8000/market' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ID": "333555",
    "Long": "-46594968",
    "Lat": "-23645941",
    "Setcens": "355030868000172",
    "Areap": "3550308005090",
    "Coddist": "69",
    "Distrito": "SACOMA",
    "CodSubPref": "13",
    "SubPrefe": "IPIRANGA",
    "Regiao5": "Sul",
    "Regiao8": "Sul 1",
    "Nome_feira": "VILA OLIVIERI",
    "Registro": "6061-5",
    "Logradouro": "RUA CLAUDIO FERREIRA MANOEL",
    "Numero": "29.000000",
    "Bairro": "VL OLIVIERI",
    "Referencia": "RUA CARLA LIVIERO"
}'
```

Alterar nova feira
```
$ curl --location --request PUT 'localhost:8000/market/2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Long": "-46594968",
    "Lat": "-23645941",
    "Setcens": "355030868000172",
    "Areap": "3550308005090",
    "Coddist": "69",
    "Distrito": "SACOMA",
    "CodSubPref": "13",
    "SubPrefe": "IPIRANGA",
    "Regiao5": "Sul",
    "Regiao8": "Sul 1",
    "Nome_feira": "VILA OLIVIERI",
    "Registro": "6061-5",
    "Logradouro": "RUA CLAUDIO FERREIRA MANOEL",
    "Numero": "29.000000",
    "Bairro": "VL OLIVIERI",
    "Referencia": "RUA CARLA LIVIERO"
}'
```
Deletar uma feira
```
$ curl --location --request DELETE 'localhost:8000/market/29'
```

# Testes
Antes de executar os testes pela primeira vez, execute o arquivo generate-mocks.sh
```
$ ./generate-mocks.sh
```
Com os mocks gerados, execute o run-tests.sh
```
$ ./run-tests.sh
```
