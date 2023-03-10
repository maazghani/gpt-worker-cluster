# Use a Go runtime image as the base image
FROM golang:alpine

# Copy the worker script into the image
COPY worker.go /worker.go

# Install the NATS client library
RUN go get github.com/nats-io/nats.go

# Set the working directory
WORKDIR /app

# Build the worker binary
RUN go build /worker.go

# Set the entrypoint to the worker binary
ENTRYPOINT ["./worker"]

