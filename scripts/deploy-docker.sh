#! /bin/bash
dos2unix backend/scripts/*
dos2unix mongo-seed/*

docker build -t localize-me-backend backend/.
docker tag localize-me-backend cafeteru/localize-me-backend:latest
docker push cafeteru/localize-me-backend:latest

docker build -t localize-me-frontend frontend/.
docker tag localize-me-frontend cafeteru/localize-me-frontend:latest
docker push cafeteru/localize-me-frontend:latest

docker build -t localize-me-spotiapp spotiapp/.
docker tag localize-me-spotiapp cafeteru/localize-me-spotiapp:latest
docker push cafeteru/localize-me-spotiapp:latest