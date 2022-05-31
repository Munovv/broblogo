FROM golang:1.18 AS build

ENV GOPATH=/
WORKDIR /src/
COPY ./ /src/

RUN go mod download; CGO_ENABLED=0 go build -o /partner-service ./cmd/main.go

# build go app
RUN go mod download
RUN go build -o partner-service ./cmd/main.go

FROM alpine:latest

# copy go app & configs
COPY --from=build /partner-service /partner-service
COPY ./configs/ /configs/

CMD ["/partner-service"]