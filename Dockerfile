# The base go-image
FROM golang:1.17-alpine

# Author and email fields
LABEL AUTHOR="ayush paharia" EMAIL="<ayush.paharia.18@gmail.com>"

# Create a directory for the app
RUN mkdir /app

# Copy all files from the current directory to the app directory
COPY . /app

COPY .env.docker /app/.env

# Set working directory
WORKDIR /app

# go build will build an executable file named server in the current directory
RUN go build -o main .

# Run the server executable
ENTRYPOINT [ "/app/main" ] 



