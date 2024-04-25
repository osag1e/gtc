# Use an official Golang runtime as a parent image
FROM golang:1.20.3-alpine

# Copy go.mod and go.sum files to the container
COPY go.mod ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

