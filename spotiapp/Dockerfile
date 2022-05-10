FROM node:lts-slim as node

WORKDIR /spotiapp

COPY . /spotiapp

RUN npm install
RUN npm run build --prod

FROM nginx:alpine

RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=node spotiapp/dist/spotiapp /usr/share/nginx/html
COPY --from=node spotiapp/nginx.conf /etc/nginx