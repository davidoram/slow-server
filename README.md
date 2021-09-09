# slow_server

A server that will delay for a requested period before returning with a textual description of how long it slept for.

To run the server:

```
go run main.go
```

To call it:

## Endpoint with a fixed delay

`GET http://localhost:8080/?d={delay}`, where `{delay}` is a [duration](https://pkg.go.dev/time#ParseDuration).

Examples:

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

## Endpoint with a random delay and failure

Will delay for a random period between 0 and 5s.

66% of the time will return `200` but 33% of the time will fail and return `500`.

Response on success is

`200` `{ "health": "ok"}` `Content-Type: application/json`

and on error is

`500` `Server error`, `Content-Type: text/plain`

```
curl "http://localhost:8080/health"

```