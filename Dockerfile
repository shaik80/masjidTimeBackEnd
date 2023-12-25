# Use the official Golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container
COPY . .

# Install make
RUN apt-get update && apt-get install -y make

# Build the Go application using the Makefile
RUN make build

# Expose the port that the application will run on
EXPOSE 4000

# Command to run the executable
CMD ["./masjid_namaz_timing", "serve"]
