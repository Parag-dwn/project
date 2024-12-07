# Step 1: Use the official Golang image for building
FROM golang:1.23.3 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Installs Go dependencies
RUN go mod download

RUN set CONFIG_PATH=./config/local.yaml  

# Step 6: Build the application
RUN go build -o main ./cmd/students-api

# Step 7: Use a minimal base image for the final build
FROM debian:bullseye-slim
# Step 8: Set the working directory inside the final image
WORKDIR /app
# Step 9: Copy the binary from the builder stage
COPY --from=builder /app/main .


# Tells Docker which network port your container listens on
EXPOSE 8085   

# Step 11: Command to run the application
CMD ["./main"]
