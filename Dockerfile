# Use the official Golang image as a parent image
FROM golang:1.21.4

# Define a build-time argument for the app name
ARG APP_NAME=trivia-backend

# Set the working directory inside the container
WORKDIR /app

# Copy the local module files to the container's workspace.
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o run . 

# Expose port 4040 to the outside
EXPOSE 8080

# Command to run the executable
CMD ["/app/run"]
