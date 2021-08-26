# slow_server

A server that will delay for a requested period before returning with a textual description of how long it slept for.

To run the server:

```
go run main.go
```

To call it:

```
# Delay 50 milliseconds
curl "http://localhost:8080/?d=50ms"
Slept for 50ms

# Delay 10 seconds
curl "http://localhost:8080/?d=10s"
Slept for 10s

# Delay 3 minutes
curl "http://localhost:8080/?d=3m"
Slept for 3m
```