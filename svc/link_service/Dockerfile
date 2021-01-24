FROM golang:latest AS builder

WORKDIR /delinkcious/link_service

ADD ./main.go main.go
ADD ./service service
ADD ./go.mod go.mod

# Fetch dependencies
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
ENV go mod tidy

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /link_service -a -tags netgo -ldflags '-s -w' .

FROM alpine:latest
COPY --from=builder /link_service /app/link_service
EXPOSE 8080
ENTRYPOINT ["/app/link_service"]
