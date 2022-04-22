# LocalizeMeFront

## Docker

- Create or update image

```shell
docker build -t localize-me-frontend .
```

- Execute Docker Image (create a container)

```shell
docker run --rm --name frontend -p 80:80 localize-me-frontend &
```

- Create tag to publish image to Docker Hub

```shell
docker tag localize-me-frontend cafeteru/localize-me-frontend:latest
```

- Publish the image to Docker Hub

```shell
docker push cafeteru/localize-me-frontend:latest
```
