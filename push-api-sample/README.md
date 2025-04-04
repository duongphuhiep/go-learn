# Push API sample

Show case how a Golang backend push notification to Frontend.

Run in Dev mode to have a more human-readable log messages in the stdout

```sh
env $(cat .env.development | xargs) go run main.go
```

Make the Backend send notification to the frontend

```sh
curl -X POST http://localhost:8080/push
```

## Advantages

Works even when the app is not open. Free, no vendor lockin.

## Disadvantages

Browser Only, Requires HTTPS, user permission, and a Service Worker.
