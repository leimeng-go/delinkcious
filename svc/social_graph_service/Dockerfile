#当前案例中没有使用go mod依赖模式，替换修改
FROM golang:latest AS builder

COPY ./../* /deliinkcious/

WORKDIR /delinkcious/svc/social_graph_service/

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /social_graph_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
MAINTAINER Ramon  <1955854667@qq.com>
COPY --from=builder /social_graph_service /app/social_graph_service
EXPOSE 9090
ENTRYPOINT ["/app/social_graph_service"]
