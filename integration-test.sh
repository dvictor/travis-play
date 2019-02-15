#!/bin/bash

set -e

sudo docker-compose up -d db
sudo docker-compose up --abort-on-container-exit --exit-code-from app app
sudo docker-compose down