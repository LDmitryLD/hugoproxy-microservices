#!/bin/bash

echo "Creating network 'my_network' if not exists..."
docker network create my_network || true

echo "Shutting down existing services..."
docker-compose -f ./proxy/docker-compose.yml down
docker-compose -f ./notify/docker-compose.yml down
docker-compose -f ./auth/docker-compose.yml down
docker-compose -f ./user/docker-compose.yml down
docker-compose -f ./geo/docker-compose.yml down

echo "Starting services..."
docker-compose -f ./geo/docker-compose.yml up --build -d
docker-compose -f ./user/docker-compose.yml up --build -d
docker-compose -f ./auth/docker-compose.yml up --build -d
echo "Waiting..."
sleep 10
docker-compose -f ./notify/docker-compose.yml up --build -d
docker-compose -f ./proxy/docker-compose.yml up --build -d

