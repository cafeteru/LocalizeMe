FROM cafeteru/node-chrome-headless as node

WORKDIR /localizeMe

COPY . /localizeMe

RUN pnpm i
RUN pnpm pre-commit
RUN pnpm build --configuration production

FROM nginx:alpine

RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=node localizeMe/dist/localize-me-front /usr/share/nginx/html
COPY --from=node localizeMe/nginx.conf /etc/nginx