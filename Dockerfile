# syntax=docker/dockerfile:1

FROM golang:1.23

# Set destination for COPY
WORKDIR /app

# Install Air
RUN go install github.com/air-verse/air@latest

# Install Lefthook
RUN go install github.com/evilmartians/lefthook@latest

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Run
# https://github.com/air-verse/air
CMD ["air", "-c", ".air.toml"]
