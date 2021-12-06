FROM golang:1.17-alpine AS BUILDER
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0
RUN mkdir -p /src
COPY . /src
WORKDIR /src
RUN go build -o /app cmd/app/main.go

FROM scratch
COPY --from=BUILDER /app /app
CMD ["/app"]
