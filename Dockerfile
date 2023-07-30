# Use the official Golang image as the base image.
FROM golang:1.18-alpine

# Set the working directory inside the container.
WORKDIR /app

# Copy the Go application source code into the container's working directory.
COPY . .

# Build the Go application inside the container.
RUN go build -o go-abs

# Expose the port that the Go application listens on.
EXPOSE 10000

# Set the command to run the Go application.
CMD ["./go-abs"]
