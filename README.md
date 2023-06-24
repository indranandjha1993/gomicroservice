# Simple Go Microservice

This is a basic microservice written in Golang. It implements a key-value store, where you can set and get values from a local map.

## Project Structure

The project follows the standard Go project layout.

```shell
/microservice
  /api
    handler.go
  /service
    service.go
  /repository
    repository.go
  /cmd
    main.go
  go.mod
  Dockerfile
```


## Getting Started

First, clone the repository to your local machine:

```bash
git clone https://github.com/indranandjha1993/gomicroservice.git
```

### Next, navigate to the project directory:

```shell
cd gomicroservice
```

### Build and run the service:
```shell
go build -o main ./cmd/main.go
./main
```

### API Endpoints
The service exposes the following endpoints:

- GET /get/{key}: Retrieve the value for a given key.
- POST /set/{key}/{value}: Set a value for a given key.

## Running with Docker
You can also build and run the service using Docker. Build the image with:

```shell
docker build -t gomicroservice .
```

### Then run the container:

```shell
docker run -p 8080:8080 gomicroservice
```

## Note
This is a very basic example and is not suitable for production use. For a real application, consider implementing proper error handling, validation, security measures, and using a real database for the repository layer.