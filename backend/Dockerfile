FROM golang:latest

WORKDIR /server

COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]