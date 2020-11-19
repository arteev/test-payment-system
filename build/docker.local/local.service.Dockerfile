
FROM golang:1.14.6-alpine3.12 as service_builder

RUN apk add --update --no-cache \
	git

ARG SERVICE
ARG LDFLAGS

WORKDIR $GOPATH/src/test-payment-system

ENV GO111MODULE=on
ENV CGO_ENABLED=0
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -ldflags="$LDFLAGS" -o=/src/app ./cmd/${SERVICE}/main.go

FROM alpine:3.12
RUN apk add --update --no-cache ca-certificates
WORKDIR /srv
# Copy a binary
COPY --from=service_builder /src/app .
ENTRYPOINT [ "./app" ]