# Greetings from fiber

Golang

`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/greetings ./src/fcgreetings`

[Image Source on Docker Hub](https://hub.docker.com/repository/docker/lucassimon/fc-greetings)

`$ dk build -t lucassimon/fc-greetings -f Dockerfile .`

`$ dk run -itd --name=fc-greetings -p 3000:3000 lucassimon/fc-greetings:latest`
