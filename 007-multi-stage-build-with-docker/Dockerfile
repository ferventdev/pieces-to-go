# Below goes the first (build) stage

# This is a standard Golang image (based on Debian Buster)
FROM golang:1.17-buster AS build

# Sets default destination for further commands
WORKDIR /build

# Downloads necessary Go modules (at least go.mod file must present)
COPY go.mod go.sum* ./
RUN go mod download

# Copies app's source files
COPY *.go ./

# Builds the artifact
RUN go build -o /minimal-server



# Below starts the second (deploy) stage

# Now we take a very lightweight base image
FROM gcr.io/distroless/base-debian10

# Only one artifact is copied into it
COPY --from=build /minimal-server /

EXPOSE 8080

# Command to execute when container is started
CMD [ "/minimal-server" ]