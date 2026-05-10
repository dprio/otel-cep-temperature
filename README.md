# cep-temperature

Este projeto é uma aplicação simples que possibilita buscar a temperatura de uma cidade a partir de um cep fornecido.

## Pré-requisitos
- Docker e Docker Compose
- Go 1.21 ou superior

## Como subir a aplicação local

`docker-compose up`

Exemplo de chamadas à API podem ser encontrados em [endpoint.http](https://github.com/dprio/otel-cep-temperature/blob/main/endpoint.http)

## Verificando traces no Zipkin

Acessar [página do Zipkin](http://localhost:9411/zipkin/?lookback=15m&endTs=1778438236306&limit=10)
