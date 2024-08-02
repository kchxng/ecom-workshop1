## Init project
- Go-Gin Web framework for RestFul API as Hex
```bash
go mod init <module-name>
```
## Download package
```bash
go mod tidy
```

## Run Application
```bash
go run .
```

## Unit Tests
```bash
go test -v
```

# Build App to Production
- Window
```bash
go build -o bin/app_server.exe . # build on windows (Recommended)
# or
go build -o bin/app_server_386.exe -a 386 main.go  # For 32-bit Windows
go build -o bin/app_server_x64.exe -a amd64 main.go  # For 64-bit Windows
```
- Linux

```bash
go build -o bin/app_server .
# or build to linux
GOOS=linux GOARCH=amd64 go build -o bin/app_server .
```

## Build Docker Container
```bash
docker compose up -d --build
docker compose down
```

## API Gateway
- traefik
- kong & konga