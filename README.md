# GolangSocks5Server
A standalone SOCKS5 server written in Go

# Building
1) Install Go (https://golang.org/doc/install) and setup your GOPATH

2) Get it
```
go get github.com/leechristensen/GolangSocks5Server
```
3) Build it
```
go install github.com/leechristensen/GolangSSHServer
```
4) Run it
```
GolangSocks5Server 2222           Starts the SOCKS server on localhost:2222

or

GolangSocks5Server 0.0.0.0:2222   Starts the SOCKS server on 0.0.0.0:2222
```

# Thanks
All credit goes to the go-socks5 project(https://github.com/armon/go-socks5)
