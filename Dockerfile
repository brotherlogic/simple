FROM golang:1.13  as build

WORKDIR /srv/grpc

COPY go.mod .
COPY proto/* ./proto/
COPY *.go ./
COPY build.sh .

ARG VERS="3.11.4"
ARG ARCH="linux-x86_64"
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v${VERS}/protoc-${VERS}-${ARCH}.zip \
      --output-document=./protoc-${VERS}-${ARCH}.zip && \
    apt update && apt install -y unzip && \
    unzip -o protoc-${VERS}-${ARCH}.zip -d protoc-${VERS}-${ARCH} && \
    mv protoc-${VERS}-${ARCH}/bin/* /usr/local/bin && \
    mv protoc-${VERS}-${ARCH}/include/* /usr/local/include

RUN CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo \
    -o /go/bin/server

FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=build /go/bin/server /server

ENTRYPOINT ["/server"]

