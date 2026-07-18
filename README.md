# Weather API Client & Cache

A performance-optimized Weather API built in Go that fetches real-time data from a third-party weather service, processes the response, and uses an in-memory Redis cache to optimize API quotas and response latency. This project demonstrates API integration, the Cache-Aside architecture pattern, data streaming transformation, and codebase modularization.

---

## Features

- **Fast Local Web Server:** Powered natively by Go's standard library (`net/http`) without extra heavy frameworks.
- **Cache-Aside Architecture:** Checks a local Redis cache before querying the external API, reducing fetch times from seconds to sub-5 milliseconds.
- **Data Transformation:** Streams massive multi-kilobyte raw JSON responses from Visual Crossing Weather API and filters out unnecessary noise to return a lightweight, tailored contract.
- **Information-Dense UI Dashboard:** Simple static HTML/JS frontend that communicates with the Go backend seamlessly over the network.
- **Secure Configuration:** Protects critical keys by utilizing an environment configuration setup via `godotenv`.
- **Modular Layout:** Codebase is cleanly organized into `models.go`, `cache.go`, `handlers.go`, and `main.go`.

---

## Technologies Used

- **Language:** Go (Golang)
- **Caching:** Redis (running via Docker)
- **Environment Management:** `github.com/joho/godotenv`
- **Redis Driver:** `github.com/redis/go-redis/v9`
- **Data Provider:** Visual Crossing Weather API

---

## Getting Started

### 1. Prerequisites

Ensure you have the following installed on your machine:

- Go 1.18 or higher
- Docker (to run the Redis instance)
- A free API key from [Visual Crossing Weather API](https://www.visualcrossing.com/)

### 2. Run Redis in Docker

Launch a background local Redis container listening on the default port `6379`:

```bash
docker run --name weather-redis -p 6379:6379 -d redis
```

### 3. Clone and Environment Setup

Create a `.env` file in the root directory of the project:

```env
WEATHER_API_KEY=your_actual_visual_crossing_api_key_here
```

### 4. Install Dependencies

Download the project dependencies specified in the modules workspace:

```bash
go mod tidy
```

### 5. Running the Server

Since the codebase is modularized across multiple companion files sharing the `main` package scope, compile and execute the entire directory using:

```bash
go run .
```

Open your browser and navigate to:

```
http://localhost:8080/
```

to test out the Weather Dashboard.

---

## API Endpoint Reference

### Get Current Weather

Returns a curated payload containing core weather indices for a target city.

**URL**

```
/weather
```

**Method**

```
GET
```

**Query Parameter**

| Parameter | Required | Description |
|-----------|----------|-------------|
| `city` | Yes | Name of the city |

**Example Request**

```http
GET http://localhost:8080/weather?city=keshod
```

**Sample JSON Response**

```json
{
  "city": "Keshod",
  "temperature": "31.0°C",
  "condition": "Partially cloudy",
  "humidity": "71%",
  "windSpeed": "22.3 km/h",
  "sunrise": "06:15:22",
  "sunset": "19:34:25",
  "icon": "partly-cloudy-day"
}
```

---

## Project Reference

This project is a complete solution for the **[Weather API Wrapper Service](https://roadmap.sh/projects/weather-api-wrapper-service)** project from **[roadmap.sh](https://roadmap.sh/)**.

---