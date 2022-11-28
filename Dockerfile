FROM golang:1.19-alpine as builder

COPY . /go/src

RUN cd src \
  && go mod download \
  && CGO_ENABLED=0 go build -v -a -ldflags "-s -w"

FROM scratch

LABEL org.opencontainers.image.title="http-dump" \
      org.opencontainers.image.version="0.0.2" \
      org.opencontainers.image.description="Dump Raw Http Requests" \
      org.opencontainers.image.licenses="Apache-2.0" \
      org.opencontainers.image.source="https://github.com/zyclonite/http-dump"

COPY --from=builder /go/src/http-dump /

EXPOSE 8080

ENTRYPOINT ["/http-dump"]
