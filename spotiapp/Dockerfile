FROM cafeteru/node-chrome-headless as node

WORKDIR /spotiapp

COPY . /spotiapp

RUN pnpm i
RUN pnpm build --configuration production

FROM nginx:alpine

RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=node spotiapp/dist/spotiapp /usr/share/nginx/html
COPY --from=node spotiapp/nginx.conf /etc/nginx