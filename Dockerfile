#LABEL Maintainer Shachindra Name shachindra@lazarus.network

FROM golang:alpine AS build-app
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o erebrus .

FROM node:18.13.0-alpine AS build-web
WORKDIR /app
COPY webapp/package*.json ./
RUN npm install
COPY webapp/ ./
RUN npm run build

FROM alpine:latest
WORKDIR /app
COPY --from=build-app /app/erebrus .
COPY --from=build-web /app/build ./webapp
COPY wg-watcher.sh .
RUN chmod +x ./erebrus ./wg-watcher.sh
RUN apk update && apk add --no-cache bash openresolv bind-tools wireguard-tools gettext inotify-tools
ENV LOAD_CONFIG_FILE=$LOAD_CONFIG_FILE RUNTYPE=$RUNTYPE SERVER=$SERVER HTTP_PORT=$HTTP_PORT GRPC_PORT=$GRPC_PORT MASTERNODE_URL=$MASTERNODE_URL
ENV REGION=$REGION DOMAIN=$DOMAIN REGION_NAME=$REGION_NAME REGION_CODE=$REGION_CODE
ENV WG_CONF_DIR=$WG_CONF_DIR WG_CLIENTS_DIR=$WG_CLIENTS_DIR WG_KEYS_DIR=$WG_KEYS_DIR WG_INTERFACE_NAME=$WG_INTERFACE_NAME
ENV WG_ENDPOINT_HOST=$WG_ENDPOINT_HOST WG_ENDPOINT_PORT=$WG_ENDPOINT_PORT WG_IPv4_SUBNET=$WG_IPv4_SUBNET WG_IPv6_SUBNET=$WG_IPv6_SUBNET
ENV WG_DNS=$WG_DNS WG_ALLOWED_IP_1=$WG_ALLOWED_IP_1 WG_ALLOWED_IP_2=$WG_ALLOWED_IP_2
ENV WG_PRE_UP=$WG_PRE_UP WG_POST_UP=$WG_POST_UP WG_PRE_DOWN=$WG_PRE_DOWN WG_POST_DOWN=$WG_POST_DOWN
ENV PASETO_EXPIRATION_IN_HOURS=$PASETO_EXPIRATION_IN_HOURS SIGNED_BY=$SIGNED_BY FOOTER=$FOOTER AUTH_EULA=$AUTH_EULA MASTERNODE_WALLET=$MASTERNODE_WALLET
RUN echo $'#!/usr/bin/env bash\n\
    set -eo pipefail\n\
    mkdir -p $WG_KEYS_DIR\n\
    /app/erebrus &\n\
    ./wg-watcher.sh\n\
    sleep infinity' > /app/start.sh && chmod +x /app/start.sh
ENTRYPOINT ["/app/start.sh"]
