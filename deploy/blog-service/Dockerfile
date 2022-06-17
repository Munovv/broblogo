FROM golang:1.18 AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

ADD ./vendor ./vendor
ADD ./go.* ./

RUN CGO_ENABLED=0 go build -o /blog-service ./cmd/blog-service/main.go

FROM alpine:latest

COPY --from=build /blog-service /blog-service
COPY ./deploy/blog-service/config.yml /configs/config.yml

CMD ["/blog-service"]






