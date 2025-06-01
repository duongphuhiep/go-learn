# datastar playground

This is a playground for [datastar] and [VictoriaLogs].

- use the [docker-compose.yaml](./docker-compose.yaml) to start a [VictoriaLogs] and a [otel-collector] instance
- then `go run main.go` to start a datastar helloworld backend at <http://localhost:8080>
- the [LoggingMiddleware](./pkg/toolspack/LoggingMiddleware.go) would log all the request / response of the backend to the [otel-collector]
- you can visualize and filter logs in [VictoriaLogs] at <http://localhost:9428/select/vmui>

## My insights

- [datastar] is very young, the framework is very tiny, so hopefully it might survive as far as the htmx! I'm using a "special" version of [datastar] (RC 11) published on their Discord channel to avoid a bug on the "indicator" plugin.

- I don't like Go app returns HTML content. Because

  - Letting Go developers to design the UI (Html, Css...) feel wrong.
  - Use the vast eco-system and tools of Node (eslint, vite..) to produce HTML, Javascript, CSS codes is better than the Go eco-system.

- [datastar] seems to be a good match for [Astro] => to be confirm. I imagine the following architecture:

  - [Astro] handles routing / static HTML, Typescript, CSS content =>
    - frontend Web developers work on this
  - [datastar] handles the client + server reactivity =>
    - Go developers work on this.
    - The same Go backend should also be re-usable for Mobile frontends.
    - Hence the backend request and response would be for "signals-only", never the HTML content.

- About testing: frontend side would implement some kind of "Mock" to stays independent of the Go backend. So that the frontend and backend developments
  would be in parallel.

## To research

- find a technology so that the datastar Go backend can describe the contract of the SSE endpoints (1 request + multiple responses)
  => candidate: [Huma] framework

[datastar]: https://data-star.dev/
[VictoriaLogs]: https://victoriametrics.com/products/victorialogs/
[otel-collector]: https://opentelemetry.io/docs/collector/
[Astro]: https://astro.build/
[Huma]: https://huma.rocks/features/server-sent-events-sse/
