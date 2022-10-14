#!/bin/bash

docker build client/ -t front
docker container run -dp 8080:8080 --name front front

docker build server/ -t back
docker container run -dp 8081:8081 --name back back

docker build db/ -t db
docker run -dp 5432:5432 --name db db -d postgres
docker exec -it db ./script.sh
