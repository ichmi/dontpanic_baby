#!/bin/bash
docker stop db
docker rm db
docker rmi db:latest
