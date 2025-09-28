# Multi-service Dockerfile
FROM docker:latest

# Install docker-compose and dependencies
RUN apk add --no-cache docker-compose

WORKDIR /app
COPY . .

# Start both services
CMD ["docker-compose", "up"]