#!/bin/bash

pg_createcluster 14 main
# service postgresql start

service postgresql stop

cat /etc/postgresql/14/main/pg_hba.conf | grep peer | sed -i -e "s/ peer/trust/" /etc/postgresql/14/main/pg_hba.conf
# service postgresql restart

service postgresql start

psql -U postgres -c "CREATE USER bexsy"
psql -U postgres -c "ALTER USER bexsy WITH ENCRYPTED PASSWORD 'frosa-ma'"
psql -U postgres -c "CREATE DATABASE daily"
psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE daily TO bexsy"

psql -U bexsy daily -c "CREATE TABLE equations (id SERIAL PRIMARY KEY, equation VARCHAR(6))"
psql -U bexsy daily -c "CREATE TABLE day_solution (id SERIAL PRIMARY KEY, solution VARCHAR(6), dt DATE)"

psql -U bexsy daily -c "INSERT INTO equations (equation) VALUES ('1+1+40'), ('2+2*20'), ('-1*-42')"
psql -U bexsy daily -c "INSERT INTO day_solution (solution, dt) VALUES ('1+1+40', NOW())"
