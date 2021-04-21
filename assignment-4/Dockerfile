# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Shubham Jagdhane <shubhamjagdhane1010@gmail.com>"

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch => 2nd stage builder
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
ENV PORT=8080 CITY=Pune REGION=PCMC COUNTRY=India IP_ADDRESS=127.0.0.1

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]
