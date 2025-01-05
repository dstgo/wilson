FROM golang:alpine AS build

#ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/build
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o gateway cmd/gateway/main.go

FROM alpine
WORKDIR /go/build

# 提供给pipeline使用
ARG APP_VERSION
ARG APP_NAME
ARG CONF_HOST
ARG CONF_TOKEN
ENV CONF_HOST=$CONF_HOST
ENV CONF_TOKEN=$CONF_TOKEN
ENV APP_NAME=$APP_NAME
ENV APP_VERSION=$APP_VERSION


COPY --from=build /go/build/gateway /go/build/gateway
CMD ["./gateway"]
