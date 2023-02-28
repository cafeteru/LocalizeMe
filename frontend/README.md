# LocalizeMeFront

## Install dependencies

```shell
npm install @angular/cli -g
npm install
```

## Run app

```shell
npm run start
```

## Commands

```shell
npm run coverage # Run the unit-tests and check the coverage of these only once
npm run coverage-watch # Run the unit-tests and check the coverage of these
npm run e2e-run # Run the e2e tests on background
npm run e2e-open # Run the e2e tests on foreground
npm run format # Format code
npm run lint # Check style, dead imports, etc...
npm run pre-commit # Perform the necessary checks before a commit
npm run test # Execute unit-test
npm run unit-test # Execute test only once
```

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

## Firebase

[Website](https://uniovi-localize-me.web.app)

Deploy:

```shell
ng build --configuration production
firebase deploy
```

### Notes

If you change the backend url, you must change its value in the environment files.
