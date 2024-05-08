FROM golang:1.22
WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build

EXPOSE 8080:8080

# Copy the backend binary
COPY ./card-jong-be .
CMD ["card-jong-be"]