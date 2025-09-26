#!/bin/bash

# Запуск
docker compose down
docker compose build
docker compose up --force-recreate -d