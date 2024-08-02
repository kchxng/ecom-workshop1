# Base image for building the go project
FROM golang:1.21-alpine AS build

# Updates the repository and installs git
RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /app

COPY . .

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/app_server .


FROM alpine:latest
RUN apk add --no-cache tzdata
ENV TZ Asia/Bangkok
COPY --from=build app_server /app/app_server
WORKDIR /app
# COPY ./authApp /app
EXPOSE 80
CMD ["/app/app_server"]