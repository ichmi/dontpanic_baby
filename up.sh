#!/bin/bash

docker compose down
docker rmi daily-back:latest daily-front:latest daily-db:latest
docker compose up
