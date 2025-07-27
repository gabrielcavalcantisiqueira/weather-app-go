# Weather App - Go

A aplicação está publicada no Google Cloud Run:

🔗 URL pública:
https://weather-app-go-418541402693.us-central1.run.app

📌 Exemplo de uso:
Consulta do clima por CEP:

```
GET /weather?cep=20021340
```

🔍 Exemplo real (clima do Rio de Janeiro):
https://weather-app-go-418541402693.us-central1.run.app/weather?cep=20021340

## Português

Este projeto é uma API simples escrita em Go que recebe um CEP, identifica a cidade e retorna a temperatura atual em graus Celsius, Fahrenheit e Kelvin.

### Objetivo

Desenvolver um sistema que:

- Receba um CEP válido (8 dígitos).
- Utilize a API ViaCEP para descobrir a cidade.
- Consulte a **WeatherAPI** com o nome da cidade para buscar a temperatura atual.
- Converta a temperatura de Celsius para Fahrenheit e Kelvin.
- Retorne os dados de forma estruturada via uma API REST.

### Tecnologias Utilizadas

- [Go](https://golang.org/)
- [ViaCEP API](https://viacep.com.br/)
- [WeatherAPI](https://www.weatherapi.com/)
- [Docker](https://www.docker.com/)


### Como executar o projeto localmente

#### 1. Clonar o repositório

```
git clone https://github.com/gabrielcavalcantisiqueira/weather-app-go.git
cd weather-app-go
```

#### 2. Definir a API Key da WeatherAPI

bash:
```
export WEATHER_API_KEY=MINHA_API_KEY
```
ou

powershell:
```
$env:WEATHER_API_KEY=MINHA_API_KEY
```

#### 3. Executar com Docker

```
docker compose up -d --build
```

## English

This project is a simple API written in Go that receives a ZIP code (CEP), identifies the city, and returns the current temperature in Celsius, Fahrenheit, and Kelvin.

### Objetivo

Develop a system that:

- Receives a valid Brazilian ZIP code (8 digits).
- Uses the ViaCEP API to find the city name.
- Queries the WeatherAPI using the city name to fetch the current temperature.
- Converts the temperature from Celsius to Fahrenheit and Kelvin.
- Returns the data in a structured format through a REST API.



### Technologies Used

- [Go](https://golang.org/)
- [ViaCEP API](https://viacep.com.br/)
- [WeatherAPI](https://www.weatherapi.com/)
- [Docker](https://www.docker.com/)

### How to run the project locally

#### 1. Clone the repository

git clone https://github.com/gabrielcavalcantisiqueira/weather-app-go.git
cd weather-app-go

#### 2. Set your WeatherAPI key

For bash:

```
export WEATHER_API_KEY=YOUR_API_KEY
```
For PowerShell:

```
$env:WEATHER_API_KEY="YOUR_API_KEY"
```

#### 3. Run with Docker

```
docker compose up -d --build
```
