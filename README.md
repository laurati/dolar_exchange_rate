# Exchange Rate

This is a Currency Quotes API.

Documentation: https://docs.awesomeapi.com.br/api-de-moedas

This API contains two servers.
The first starts on port :8080 and uses the POSTGRESQL database to save data.
The second starts on port :8082 and uses the SQLITE database to save data.

## Functionalities

- Query exchange rate by code and save it
- Query exchange rates
- Query dolar exchange rate

## Prerequisites

- Docker
- Docker-compose

## Installation

Clone this repository:
```shell
git clone https://github.com/laurati/exchange_rate.git
```

## Running Bundles Application
1. Navigate to the application repository:
```shell
cd exchange_rate
```

2. Inside the repository in root, execute the command:
```shell
docker-compose up -d
```

3. Inside the repository in root, execute the command:
```shell
go run cmd/api/main.go
```
or

```shell
go run dolar/main.go
```
