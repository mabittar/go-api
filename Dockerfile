# Use the official Go image
FROM golang:1.24-alpine

# Set the working directory
WORKDIR /app

# Copy Go modules and dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main .

# Expose the API port
EXPOSE 8080

# Run the application
CMD ["./main"]