# Build stage
FROM golang:1.24-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy Go modules and dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main .

# Final stage
FROM scratch

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/main .

# Expose the API port
EXPOSE 8080

# Run the application
CMD ["./main"]