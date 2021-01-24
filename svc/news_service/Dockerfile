FROM golang:latest AS builder

WORKDIR /delinkcious/news_service

ADD ./main.go main.go
ADD ./service service
ADD ./go.mod go.mod

# Fetch dependencies
# ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
RUN go mod tidy

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /news_service -a -tags netgo -ldflags '-s -w' .

FROM alpine:latest
COPY --from=builder /news_service /app/news_service
EXPOSE 7070
ENTRYPOINT ["/app/news_service"]
