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

token
```text
token 通过向 telegram @BotFather 注册机器人后生成.
```

socks5
```text
无法访问 telegram 的网络需要使用socks5代理进行 telegram 客户端连接.
```
