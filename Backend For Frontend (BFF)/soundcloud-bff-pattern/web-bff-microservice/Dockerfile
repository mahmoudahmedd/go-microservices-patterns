# Use the official Go image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /web-bff-microservice

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o web-bff-microservice .

# Set the HOST and PORT environment variables
ENV PORT=8080

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./web-bff-microservice"]