# Start from golang base image
FROM golang:alpine as builder

ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Shubham Jagdhane <shubhamjagdhane1010@gmail.com>"

# Set the current working directory inside the container 
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# install dependencies
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage container => 2nd stage builder
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
ENV PORT=8080 CITY=Pune REGION=PCMC COUNTRY=India IP_ADDRESS=127.0.0.1

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]
