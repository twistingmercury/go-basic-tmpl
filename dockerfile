FROM golang:1.21-alpine3.19 as build

ENV CGO_ENABLED=0 GO_ENABLED=0

ARG BUILD_DATE
ARG BUILD_VER

RUN apk --no-cache add ca-certificates

WORKDIR /src
COPY . .

RUN apk add --no-cache git && \
    go mod tidy && \
    go build -ldflags \
    "-X 'MODULE_NAME/conf.buildDate=${BUILD_DATE}' \
     -X 'MODULE_NAME/conf.buildVer=${BUILD_VER}' -s -w" \
    -o ./bin/BIN_NAME main.go

FROM scratch as final

ARG BUILD_DATE
ARG BUILD_VER
ARG GIT_COMMIT

LABEL \
    org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.version="${BUILD_VER}" \
    org.opencontainers.image.vendor="IMG_VENDOR_NAME" \
    org.opencontainers.image.description="IMG_DESCRIPTION" \
    org.opencontainers.image.title="BIN_NAME" \
    com.mcg.container.help="docker exec -it <CONTAINER> /app/BIN_NAME --help" \
    com.mcg.container.os="scratch"

WORKDIR /app
COPY --from=build /src/bin/ /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/BIN_NAME"]
