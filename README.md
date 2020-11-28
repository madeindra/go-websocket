# Websocket in Go

Tutorial by [Tabvn](https://www.youtube.com/watch?v=yyREnTgRTQ0)


## How to run websocket

```
go run main.go
```

## How to start chat
Go to `localhost:3000`

### Subs
```
ws.send('{"action": "subscribe", "topic": "abc"}')
```

### Pub
```
ws.send('{"action": "publish", "topic": "abc", "message":"Hi There!"}')
```
