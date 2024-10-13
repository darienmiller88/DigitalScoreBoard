# Build stage
FROM golang:1.23.1-alpine AS build

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o deploy .

# Final stage (runtime)
FROM alpine:latest
WORKDIR /app

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy the Go binary, templates and static folder from build stage
COPY --from=build /app/deploy /deploy

# Ensure the binary is executable
RUN chmod +x /deploy

# Expose the port
EXPOSE ${PORT}

# CMD to start the Go app
CMD ["/deploy"]
