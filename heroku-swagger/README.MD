# Heroku-Swagger

## Online version

[Website](https://localize-me-open-api.herokuapp.com)

## How to update swagger

- Convert `..backend/swagger/swagger.json` to `swagger.yml`
- Update `dist/swagger.yml` with this content.
- Execute this commands if the repository isn´t init

```shell
git init
heroku git:remote -a <project-name>
```

- Update commit

```shell
git add .
git commit -am <commit-message>
git push heroku master
```
