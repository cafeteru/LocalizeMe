#! /bin/bash
dos2unix.exe backend/scripts/*
dos2unix.exe mongo-seed/*

docker build -t localize-me-backend backend/.
docker tag localize-me-backend igm1990/localize-me-backend:latest
docker push igm1990/localize-me-backend:latest

docker build -t localize-me-frontend frontend/.
docker tag localize-me-frontend igm1990/localize-me-frontend:latest
docker push igm1990/localize-me-frontend:latest