FROM golang:latest
LABEL maintainer="Joris van de Donk<jorisvddonk@gmail.com"
WORKDIR /app

# Get dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source files
COPY . .

RUN ls -la

# Build app
RUN go build -o main .

CMD ["./main"]