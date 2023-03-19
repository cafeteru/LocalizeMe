#! /bin/bash
dos2unix.exe backend/scripts/*
docker compose up -d
docker rm localize-me-mongo-seed