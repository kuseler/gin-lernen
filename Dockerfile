<<<<<<< HEAD
# Start from the official Golang image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . ./

# Build the Go application
RUN go build -o main .

# Expose the application's port
EXPOSE 9999

# Set the entrypoint command
CMD ["./main"]
=======
# Start from the official Golang image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . ./

# Build the Go application
RUN go build -o main .

# Expose the application's port
EXPOSE 9999

# Set the entrypoint command
CMD ["./main"]
>>>>>>> 43dfc5aefbd0f2da92f9237fb13c126e96d4f5c1
