FROM node:lts-slim as node

WORKDIR /localizeMe

COPY . /localizeMe

RUN npm install
RUN npm run build --prod


FROM nginx:alpine

COPY --from=node localizeMe/dist/localize-me-front /usr/share/nginx/html
COPY --from=node localizeMe/conf.d /etc/nginx/conf.d