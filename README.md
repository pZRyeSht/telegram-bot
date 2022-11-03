# Telegram_Bot
Implementation of telegram bot

```go
windows build mac
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
```

```go
mac build windows or linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

```go
linux build mac or windows
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

usage
```shell
./telegram-bot -token=12345678901:XXXXXWM9rDuKEiTjLahVcuYVK5lGKLLV-00 -socks5=socks5://localhost:1090
```
