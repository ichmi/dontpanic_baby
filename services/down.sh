#!/bin/bash

docker stop db front back
docker rm db front back
docker rmi db:latest front:latest back:latest
