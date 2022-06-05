# Spotiapp

## Install dependecies

```shell
npm install
```

## Run app

```shell
npm run start
```

## Docker

- Create or update image

```shell
docker build -t localize-me-spotiapp .
```

- Execute Docker Image (create a container)

```shell
docker run --rm --name spotiapp -p 80:80 localize-me-spotiapp &
```

- Create tag to publish image to Docker Hub

```shell
docker tag localize-me-spotiapp cafeteru/localize-me-spotiapp:latest
```

- Publish the image to Docker Hub

```shell
docker push cafeteru/localize-me-spotiapp:latest
```

### Notes

If you change the backend url or spotify credentials, you must to change its value in the environment files.
