FROM golang:1.13-alpine AS build

WORKDIR /src/
COPY . .
RUN go mod download && \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/greetings ./src/fcgreetings

FROM scratch
USER 1000
WORKDIR /app/
COPY --from=build /src/build/greetings /app/greetings
COPY /src/fcgreetings/views /app/views

EXPOSE 3000

ENTRYPOINT ["/app/greetings"]
