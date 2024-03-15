# Base image
FROM golang:1.22.1-bookworm AS Builder

# Set the working directory inside the container
WORKDIR /app

# Copy the application source code
COPY . .

# Download and cache Go modules
RUN go mod download

# Build the application
RUN go build -o ./appbin ./main.go

FROM debian:bookworm

# Set the working directory inside the container
WORKDIR /app

# Copy the application binary from the builder image
COPY --from=Builder /app/appbin ./
COPY --from=Builder /app/assets ./assets

# Download FFMPEG
RUN apt-get update
RUN apt-get install -y ffmpeg

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./appbin"]
