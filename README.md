# UniOVI-LocalizeMe

Real-time management system for localization strings

## How to start all applications

### Create Containers

```shell
docker compose up # run in the foreground showing logs
docker compose up -d # execute at background
```

After docker-compose is complete, applications will be running on the following ports:

- Backend: 8080.
- Swagger (Backend's OpenApi): 9090.
- Frontend: 80.
- SpotiApp (Sample application using LocalizeMe): 90.
- MongoDb (database): 27017.

## How to finish all applications

```shell
docker compose down
```

### Update Docker's images

Execute `deploy-docker.sh` to update Docker's images.
