# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to your app's directory
WORKDIR /go/src/microservices-test

# Copy the go.mod and go.sum files from the parent directory
COPY ../go.mod ./
COPY ../go.sum ./

RUN  go mod download

COPY . ./

ENV CGO_ENABLED=0
ENV GOOS=linux

# Build the Helper service
RUN go build -o book/main ./book

# Expose the port the service runs on
EXPOSE 12000

# Run the Helper service
CMD ["./book/main"]
