FROM golang:1.24.6-alpine AS buildbase

RUN apk add git build-base ca-certificates
WORKDIR /go/src/github.com/deantech-io/noaa-indexer
COPY . .

RUN CGO_ENABLED=1 GO111MODULE=on GOOS=linux go build -o /usr/local/bin/noaa-indexer /go/src/github.com/deantech-io/noaa-indexer

# Local development only (see compose.yaml).
FROM buildbase AS mockrelayer-build
RUN CGO_ENABLED=1 GO111MODULE=on GOOS=linux go build -o /usr/local/bin/mockrelayer /go/src/github.com/deantech-io/noaa-indexer/devtools/mockrelayer

FROM scratch AS mockrelayer
COPY --from=alpine:3.22 /bin/sh /bin/sh
COPY --from=alpine:3.22 /usr /usr
COPY --from=alpine:3.22 /lib /lib

COPY --from=mockrelayer-build /usr/local/bin/mockrelayer /usr/local/bin/mockrelayer

ENTRYPOINT ["mockrelayer"]

FROM scratch
COPY --from=alpine:3.22 /bin/sh /bin/sh
COPY --from=alpine:3.22 /usr /usr
COPY --from=alpine:3.22 /lib /lib

COPY --from=buildbase /usr/local/bin/noaa-indexer /usr/local/bin/noaa-indexer
COPY --from=buildbase /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["noaa-indexer"]
