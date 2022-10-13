FROM golang:1.19.2 AS BUILD

ENV VERSION 0.1.1

# build
ADD . /build
WORKDIR /build
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o yojee

FROM alpine
LABEL maintainer="yuzu <like09th@gmail.com>"

COPY --from=BUILD /build/yojee /app/yojee

CMD ["/app/yojee"]