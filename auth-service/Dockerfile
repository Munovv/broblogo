FROM golang:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

RUN go version
ENV GOPATH=/

COPY ./ ./

# build auth-service
RUN go mod download
RUN go build -o auth-service ./cmd/main/main.go

CMD ["./auth-service"]