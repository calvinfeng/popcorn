# API

## Docker

    docker build -t api .
    docker run -p 8000:8000 api

## Server

### Endpoints

```text
GET api/recommendations
GET api/auth
```

## UI (Progressive Web App)

For development, please use `lvh.me:8000` instead of `localhost:8000`.