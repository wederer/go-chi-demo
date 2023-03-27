FROM golang:1.20 AS build

WORKDIR /app

# manage dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o /app-build /app/cmd/web/main.go

FROM alpine:latest
WORKDIR /
COPY --from=build /app-build /app-build
CMD /app-build
