#当前案例中没有使用go mod依赖模式，替换修改
FROM golang:latest AS builder

WORKDIR /delinkcious/social_graph_service

ADD ./main.go main.go
ADD ./service service
ADD ./go.mod go.mod
#Fetch dependencies
# ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
RUN go mod tidy

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /social_graph_service -a -tags netgo -ldflags '-s -w' .

FROM alpine
COPY --from=builder /social_graph_service /app/social_graph_service
EXPOSE 9090
ENTRYPOINT ["/app/social_graph_service"]
