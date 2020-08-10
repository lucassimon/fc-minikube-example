FROM golang:1.13-alpine AS build

WORKDIR /src/
COPY /build/greetings /bin/greetings

FROM scratch
WORKDIR /opt/src/
COPY --from=build /bin/greetings /opt/src/greetings
COPY /src/fcgreetings/views /opt/src/views

EXPOSE 3000

ENTRYPOINT ["/opt/src/greetings"]
