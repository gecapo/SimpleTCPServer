# SimpleTCPServer
Simple tcp server written in go

## Usage

To start the server

```
go run main.go
```
 or you can use optional parameters : port, host, conn
 
```
go run main.go -port=3333 -host=localhost
```

To send messages to the server (default port is 6969)

```
echo -n "some message" | nc localhost 6969
```

You should get a response like 
```
"Message received."
```
