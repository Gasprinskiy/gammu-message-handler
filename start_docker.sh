#!/bin/bash

# Запуск
docker compose down
docker compose build
docker compose -p tgsms up --force-recreate --remove-orphans