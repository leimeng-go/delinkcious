FROM golang:latest AS builder
ADD ./main.go main.go
ADD ./service service
# Fetch dependencies
RUN go mod tidy

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /link_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
MAINTAINER Ramon <1955854667@qq.com>
COPY --from=builder /link_service /app/link_service
EXPOSE 8080
ENTRYPOINT ["/app/link_service"]
