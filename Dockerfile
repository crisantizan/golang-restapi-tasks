FROM golang:1.14.4-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o golang-restapi-tasks -v .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /bin

# Copy binary from build to main folder
RUN cp /build/golang-restapi-tasks .

# Build a small image
FROM scratch

COPY --from=builder /bin/golang-restapi-tasks /
COPY ./data.json /data.json

# Command to run
ENTRYPOINT ["/golang-restapi-tasks"]