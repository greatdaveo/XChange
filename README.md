# XChange is a real-time currency converter API.

**XChange** is a currency conversion API built with Go. It fetches real-time exchange rates from multiple external providers and offers accurate conversion between currencies. The service is Dockerized for easy deployment and consistent environments.

---

## Features

* Real-time currency conversion between any two supported currencies
* Fallback system using multiple exchange rate providers: 

  * ExchangeRate API
  * CurrencyLayer API
* Clear separation of concerns using interfaces and services
* Environment variable support via `.env` for API keys
* Containerized using Docker for easy setup and deployment

---

## How It Works

### `/convert` endpoint

Example usage:

```
GET /convert?amount=100&from=USD&to=NGN
```

The service performs the following:

1. Attempts to fetch the exchange rate using the ExchangeRate API
2. If that fails, it falls back to the CurrencyLayer API
3. Returns the converted amount, rate used, and a timestamp

---

## Running Locally with Docker

### 1. Set up your `.env` file

Create a `.env` file in the root directory:

```
EXCHANGE_RATE_API_KEY=your_key_here
CURRENCYLAYER_API_KEY=your_key_here
```

### 2. Build the Docker image

```bash
docker build -t xchange-app .
```

### 3. Run the container

```bash
docker run -p 8080:8080 --env-file .env xchange-app
```

### 4. Test the endpoint

Open your browser or use curl/Postman:

```
http://localhost:8080/convert?amount=100&from=USD&to=NGN
```

---

## Example JSON Response

```json
{
  "from": "USD",
  "to": "NGN",
  "original_amount": 100,
  "converted_amount": 152777.65,
  "rate_used": 1527.7765,
  "retrieved_at": "2025-07-17T23:11:30Z"
}
```
---
---

## Developed By
> Olowomeye David [GitHub](https://github.com/greatdaveo) [LinkedIn](https://linkedin.com/in/greatdaveo)

---
