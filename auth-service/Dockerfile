FROM golang:1.18 AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

RUN go mod download; CGO_ENABLED=0 go build -o /auth-service ./cmd/main/main.go

# build go app
RUN go mod download
RUN go build -o auth-service ./cmd/main/main.go

FROM alpine:latest

# copy go app & configs
COPY --from=build /auth-service /auth-service
COPY ./configs/ /configs/

CMD ["/auth-service"]