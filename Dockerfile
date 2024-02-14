FROM golang:alpine AS build
WORKDIR /usr/src/app
COPY go.* ./
RUN go mod tidy
COPY . .
RUN GOOS=linux go build -o rinha ./cmd/main.go
FROM alpine:latest
COPY --from=build /usr/src/app/rinha /usr/local/bin
EXPOSE 8080
CMD [ "/usr/local/bin/rinha" ]
