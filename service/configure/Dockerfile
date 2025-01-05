# 构建go
FROM golang:alpine AS gobuild
#ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/build
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o configure cmd/configure/main.go


FROM alpine
WORKDIR /go/build
COPY ./internal/conf/conf-prod.yaml /go/build/internal/conf/conf.yaml
COPY ./static /go/build/static
COPY ./deploy /go/build/deploy
COPY ./web /go/build/web
COPY --from=gobuild /go/build/configure /go/build/configure
CMD ["./configure"]
