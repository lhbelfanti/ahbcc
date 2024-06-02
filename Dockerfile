FROM golang:alpine

LABEL maintainer="Lucas Belfanti"

# Install necessary dependencies
RUN apk update && \
    apk add --no-cache git bash build-base

# Create the application directory and set it as the working directory
WORKDIR /app

# Copy only the Go module files to leverage caching and download go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code the migrations folder and the .env configuration file
COPY cmd/ ./cmd
COPY internal/ ./internal
COPY migrations/ ./migrations
COPY .env ./

# Build the application and output the binary as 'build'
RUN CGO_ENABLED=0 GOOS=linux go build -o /ahbcc ./cmd

# Expose port
EXPOSE 8080

# Run application
CMD [ "/ahbcc" ]

