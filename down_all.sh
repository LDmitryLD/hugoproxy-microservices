#!/bin/bash
echo "Shutting down existing services..."
docker-compose -f ./proxy/docker-compose.yml down
docker-compose -f ./notify/docker-compose.yml down
docker-compose -f ./auth/docker-compose.yml down
docker-compose -f ./user/docker-compose.yml down
docker-compose -f ./geo/docker-compose.yml down