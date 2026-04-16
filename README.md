# URL Shortener

A simple URL shortening service built with Go. Convert long URLs into compact, shareable short codes with analytics tracking.

## 🚀 Features

- **URL Shortening** - Convert long URLs into short, memorable codes
- **Click Analytics** - Track clicks for each shortened URL
- **URL Validation** - Protocol and format validation to prevent invalid URLs
- **Distributed ID Generation** - Snowflake-based unique ID generation for scalability
- **Environment Configuration** - Configurable via `.env` file
- **RESTful API** - Clean, intuitive API endpoints
- **Database Migrations** - Automatic schema setup on startup

## 📋 Prerequisites

- Go 1.16+
- PostgreSQL 12+

## 🛠️ Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/srivatsa17/url-shortener.git
   cd url-shortener
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Set up environment variables**

   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Build the application**
   ```bash
   make build
   ```

## 📚 Usage

### Start the server

```bash
make run
```

### API Endpoints

#### 1. Shorten URL

**POST** `/api/v1/shorten`

Request:

```json
{
  "url": "https://example.com/very/long/url/path"
}
```

Response:

```json
{
  "long_url": "https://example.com/very/long/url/path",
  "short_url": "http://localhost:8080/r/abc123",
  "created_at": "2026-04-16T10:30:45Z",
  "click_count": 0
}
```

#### 2. Redirect to Original URL

**GET** `/api/v1/redirect?code=abc123`

- Redirects to the original long URL
- Increments click counter automatically
- Returns 404 if short code not found

#### 3. Health Check

**GET** `/health`

Response:

```json
{
  "status": "healthy"
}
```

## 🧪 Testing

```bash
make test
```
