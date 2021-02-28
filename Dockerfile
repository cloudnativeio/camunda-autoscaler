FROM golang:1.14.10-alpine AS builder

WORKDIR /go/src

ENV GO111MODULE=on
ENV GOFLAGS -mod=vendor
ENV GCO_ENABLED=0

COPY ./ /go/src/

RUN go build -v -o autoscaler main.go

FROM alpine:3.12.3

COPY --from=builder /go/src/autoscaler /app/

WORKDIR /app

CMD [ "/app/autoscaler" ]