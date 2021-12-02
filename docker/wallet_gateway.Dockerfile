FROM golang:1.17-alpine AS BUILDER
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0
RUN mkdir -p /src
COPY . /src
WORKDIR /src
RUN go build -o /gateway cmd/wallet_gateway/main.go

FROM scratch
COPY --from=BUILDER /gateway /gateway
CMD ["/gateway"]
