FROM golang:1.18 AS builder

WORKDIR /build

ADD ./go.mod  go.mod
ADD ./main.go main.go
ADD ./service service

# Update
RUN apt-get --allow-releaseinfo-change update && apt upgrade -y

# Fetch dependencies
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/ && go mod download all

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /social_graph_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
MAINTAINER Gigi Sayfan <the.ttlzx@gmail.com>
COPY --from=builder /social_graph_service /app/social_graph_service
EXPOSE 9090
ENTRYPOINT ["/app/social_graph_service"]
