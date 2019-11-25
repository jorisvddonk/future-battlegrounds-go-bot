# Future Battlegrounds - Go Bot (WIP)

This is a simple example bot for Future Battlegrounds, written in Go, using the GRPC API

## Current status

Currently, the only thing this does is connect with the Future Battlegrounds server at localhost, spawn a ship, and log the battleground state!

## Regenerate protobuf client

```
protoc -I ../future-battlegrounds/src/main/proto futurebattlegrounds.proto --go_out=plugins=grpc:futurebattlegrounds
```

## Running

```
go run main.go
```
