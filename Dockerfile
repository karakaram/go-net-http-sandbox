# dev, builder
FROM golang:1.14.2 AS golang
WORKDIR /work/go-web-application-sandbox
ENV GO111MODULE=on

# dev
FROM golang as dev
RUN curl -fsLo /usr/local/bin/air https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air \
  && chmod +x /usr/local/bin/air

# builder
FROM golang AS builder
COPY ./ ./
RUN make prepare build-linux

# release
FROM alpine AS app
COPY --from=builder /work/go-web-application-sandbox/build/go-web-application-sandbox-linux-amd64 /usr/local/bin/go-web-application-sandbox
EXPOSE 8080
ENTRYPOINT ["go-web-application-sandbox"]
