# syntax=docker/dockerfile:1
FROM golang:1.20-alpine
WORKDIR /app

# Copy the current directory contents into the container at /app

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy
COPY . .

# Build the Golang application
RUN go build .

# Expose port 8080
EXPOSE 4001

# Set the entrypoint to run the Golang application
CMD ["go", "run", "main.go"]
