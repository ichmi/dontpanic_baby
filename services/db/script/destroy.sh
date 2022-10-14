#!/bin/bash

dropdb -U postgres daily --force --if-exists
dropuser -U postgres bexsy --if-exists
