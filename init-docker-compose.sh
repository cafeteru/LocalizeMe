#! /bin/bash
dos2unix.exe backend/scripts/*
dos2unix.exe mongo-seed/*
docker compose up -d