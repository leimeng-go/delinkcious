module github.com/pingguodeli573365/delinkcious

go 1.15

replace github.com/nats-io/go-nats v1.10.0 => github.com/nats-io/nats.go v1.10.0

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/go-kit/kit v0.10.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/lib/pq v1.9.0
	github.com/nats-io/go-nats v1.10.0
	github.com/nuclio/nuclio-sdk-go v0.2.0
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.4
	github.com/pelletier/go-toml v1.8.1
	github.com/prometheus/client_golang v1.9.0
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	google.golang.org/grpc v1.35.0
	gopkg.in/yaml.v2 v2.4.0
)
