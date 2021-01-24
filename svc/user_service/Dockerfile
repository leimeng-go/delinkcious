FROM golang:latest AS builder

WORKDIR /delinkcious/user_service

ADD ./main.go main.go
ADD ./service service
ADD ./go.mod go.mod
# Fetch dependencies
# ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
RUN go mod tidy

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /user_service -a -tags netgo -ldflags '-s -w' .

FROM alpine:latest
COPY --from=builder /user_service /app/user_service
EXPOSE 7070
ENTRYPOINT ["/app/user_service"]
