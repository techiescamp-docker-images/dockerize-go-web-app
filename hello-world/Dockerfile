# Stage 1: Build the Go application
FROM golang:1.20-alpine AS build

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o server .

# Stage 2: Create a minimal image for running the application
FROM alpine:3.18

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /app

# Copy the built application from the previous stage
COPY --from=build /app/server /app/server
COPY --from=build /app/index.html /app/index.html

# Set the appropriate permissions
RUN chown -R appuser:appgroup /app && \
    chmod +x /app/server

# Switch to the non-root user
USER appuser

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./server"]
