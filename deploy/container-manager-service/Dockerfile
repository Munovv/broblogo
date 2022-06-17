FROM golang:1.18 AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

ADD ./vendor ./vendor
ADD ./go.* ./

RUN CGO_ENABLED=0 go build -o /container-manager-service ./cmd/container-manager-service/main.go

FROM alpine:latest

RUN apk update && \
    apk add --no-cache docker-cli python3 && \
    apk add --no-cache --virtual .docker-compose-deps python3-dev libffi-dev openssl-dev gcc libc-dev make && \
    pip3 install docker-compose && \
    apk del .docker-compose-deps

# copy go app & configs
COPY --from=build /container-manager-service /container-manager-service
COPY ./deploy/container-manager-service/config.yml /configs/config.yml
COPY ./deploy/docker-compose.yml /compose/docker-compose.yml

CMD ["/container-manager-service"]