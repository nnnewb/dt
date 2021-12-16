FROM golang:1.17-alpine AS BUILDER
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0
RUN mkdir -p /src
COPY . /src
WORKDIR /src
RUN go build -o /tm cmd/tm/main.go

FROM scratch
COPY --from=BUILDER /tm /tm
CMD ["/tm"]
