FROM node:lts-slim as node

WORKDIR /localizeMe

COPY . /localizeMe

RUN npm install
RUN npm run build --prod


FROM nginx:alpine

RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=node localizeMe/dist/localize-me-front /usr/share/nginx/html
COPY --from=node localizeMe/nginx.conf /etc/nginx