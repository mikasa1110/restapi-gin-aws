# Use a Go base image
FROM golang:1.19
# Set the GOPROXY environment variable
ENV GOPROXY=https://goproxy.io,direct
# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Install Go module dependencies
RUN go mod download

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go app
RUN go build -o main

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["/app/main"]