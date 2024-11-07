FROM golang:1.23 AS builder

WORKDIR /go/src
COPY ./ ./

# build
RUN make build WORKSPACE=srv-kubelie

# runtime
FROM alpine:latest

ARG PROJECT_NAME=srv-kubelie

COPY --from=builder /go/src/cmd/${PROJECT_NAME}/${PROJECT_NAME} /go/bin/${PROJECT_NAME}

EXPOSE 80

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/srv-kubelie"]