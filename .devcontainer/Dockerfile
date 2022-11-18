FROM harbor.avlyun.org/inf/golang:1.17 as goBuilder
WORKDIR /go_build
COPY . /go_build
ENV GOPROXY=https://registry-mirrors.avlyun.org/repository/goproxy/,https://goproxy.io,direct

RUN go build -mod=vendor -ldflags "-X main.Version=$VERSION_TAG" -o /go/bin/ ./...
# debian release as the same as golang image
# set TimeZone as Asia/Shanghai
# set Local as zh-hans
FROM debian:stable
EXPOSE 8181
COPY --from=goBuilder /go/bin/* .
VOLUME /records
COPY ./records /records

ENTRYPOINT ["./file_sort"]