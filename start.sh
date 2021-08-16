#!/bin/bash
DB_USER=admin \
DB_PASSWD=123456 \
DB_NAME=unico \
FILE_NAME=DEINFO_AB_FEIRASLIVRES_2014 \

mysql -u$DB_USER -p$DB_PASSWD < init.sql

mysqlimport --ignore-lines=1 --lines-terminated-by='\n' --fields-terminated-by=',' --fields-enclosed-by='"' --verbose --local -u$DB_USER -p$DB_PASSWD $DB_NAME $FILE_NAME.csv

go run main.go