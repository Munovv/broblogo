FROM golang:1.18 AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

ADD ./vendor ./vendor
ADD ./go.* ./

RUN CGO_ENABLED=0 go build -o /partner-service ./cmd/partner-service/main.go

FROM alpine:latest

# copy go app & configs
COPY --from=build /partner-service /partner-service
COPY ./deploy/partner-service/config.yml /configs/config.yml

CMD ["/partner-service"]