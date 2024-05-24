# Use the official Golang image as a base image
FROM golang:1.22.3-bullseye

# Set the working directory inside the container
WORKDIR /app

# Download and install dependencies
# RUN go mod download

# Copy the entire project to the container
COPY . .

# Build the Go application
RUN go build -o /app/main ./cmd/main.go

# Expose the port the application will run on
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]