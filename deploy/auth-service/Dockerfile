FROM golang:1.18 AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

ADD ./vendor ./vendor
ADD ./go.* ./

RUN CGO_ENABLED=0 go build -o /auth-service ./cmd/auth-service/main.go

FROM alpine:latest

# copy go app & configs
COPY --from=build /auth-service /auth-service
COPY ./deploy/auth-service/config.yml /configs/config.yml

CMD ["/auth-service"]