# Simple RabbitMQ with golang

## Build Info ⚙
You can create a docker container with the floowing command.
```
build rabbit mq container docker run -d --hostname rabbitmq --name test-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```
Navigate to localhost:15672 and use username:guest and password:guest to access rabbitmq management web app.

## Project structure 📁
- `rabbitmq.go` Logic for rabbit mq
- `main.go` 