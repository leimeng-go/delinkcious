FROM golang:latest AS builder
ADD ./main.go main.go
ADD ./service service
# Fetch dependencies
RUN go mod tidy

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /news_service -a -tags netgo -ldflags '-s -w' .
FROM scratch
MAINTAINER Ramon <1955854667@qq.com>
COPY --from=builder /news_service /app/news_service
EXPOSE 7070
ENTRYPOINT ["/app/news_service"]
