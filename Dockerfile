# Use the official Go image as the build environment
FROM golang:1.22 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download all dependencies. This will be cached if go.mod and go.sum don't change
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN make build

# Start a new stage with a minimal base image
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/out/kickstart /kickstart

# Specify the entrypoint for the container (optional, not needed for CLI extraction)
ENTRYPOINT ["/kickstart"]
