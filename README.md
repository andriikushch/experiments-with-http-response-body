# Experiments with http response body

This repository contains code examples used in the article about what can go wrong if http.Response Body is handled in the wrong way.

## Start the server

To build and start:
```
cd server
go build server.go
./server
```

In the output, you will see the PID of the server process.

## Start the client

To build 
```
cd client
go build client.go
```

Client has two parameters:

```
  -number int
        specify the number of requests to send (default 5)
  -type string
        specify type of request to send:
        close                   - send and close
        read                    - send and read
        readandclose    - send, read and close
        nothing                 - send
         (default "readandclose")
```

Example:

`./client --number 1000 --type nothing`

In the output, you will see the PID of the client process.

## In case of an error

In case of an error, please read the error message. 

It might be that port 8080 can be already used by another process. In this case, please change it in the source code and repeat the build step.