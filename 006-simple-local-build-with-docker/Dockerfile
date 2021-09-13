FROM golang:1.17-alpine

# Sets default destination for further commands
WORKDIR /build

# Downloads necessary Go modules (at least go.mod file must present)
COPY go.mod go.sum* ./
RUN go mod download

# Copies app's source files
COPY *.go ./

# Compiles the program
RUN go build -o /minimal-server

EXPOSE 8080

# Command to execute when container is started
CMD [ "/minimal-server" ]