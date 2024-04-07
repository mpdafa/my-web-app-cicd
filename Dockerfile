FROM golang:1.21 as builder

# Create and change to the app directory.
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o my-web-app-cicd

FROM debian:buster-slim
WORKDIR /root/

COPY --from=builder /app/my-web-app-cicd .

# Install necessary packages (if any).
# For example, if your app needs to make HTTPS requests, SSL certificates might be needed.
# RUN apt-get update && apt-get install -y \
 #   ca-certificates \
 #   && rm -rf /var/lib/apt/lists/*

# Expose port 8888 to the outside world.
EXPOSE 8888

# Run the web service on container startup.
CMD ["./my-web-app-cicd"]