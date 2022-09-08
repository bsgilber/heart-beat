# stage 1: build binary
FROM golang:1.17-buster as build

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# build the binary
RUN make build

# stage 2: launch binary
FROM alpine:3.14 AS run

# Create and change to the app directory.
WORKDIR /app

# Copy binary from the build layer to this layer
COPY --from=build /app/dist/main .

ENV ENVIRONMENT="local"

EXPOSE 8080
CMD ["/app/main"]

