FROM golang:1.14.10-alpine AS builder

WORKDIR /go/src

ENV GO111MODULE=on
ENV GOFLAGS -mod=vendor
ENV GCO_ENABLED=0

COPY ./ /go/src/

RUN go build -v -o autoscaler main.go

FROM hub.artifactory.gcp.anz/alpine:3.12.3

COPY --from=builder /go/src/ob-vault-gopher /app/

WORKDIR /app

CMD [ "/app/autoscaler" ]