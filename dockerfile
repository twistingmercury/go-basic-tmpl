ARG GO_VERSION=1.22
ARG ALPINE_VERSION=3.19

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build

ENV CGO_ENABLED=0 GO_ENABLED=0

ARG BUILD_DATE
ARG BUILD_VER
ARG GIT_COMMIT
ARG BIN_NAME
ARG MODULE_NAME
ARG TARGET

RUN apk --no-cache add ca-certificates

WORKDIR /src
COPY . .

RUN apk add --no-cache git && \
    go mod tidy && \
    go build -ldflags \
    "-X '$MODULE_NAME/cmd/conf.buildDate=${BUILD_DATE}' \
     -X '$MODULE_NAME/cmd/conf.buildVer=${BUILD_VER}' \
     -X '$MODULE_NAME/cmd/conf.buildCommit=${GIT_COMMIT}' -s -w" \
    -o ./bin/${BIN_NAME} ${TARGET}

FROM scratch as final

ARG BUILD_DATE
ARG BUILD_VER
ARG GIT_COMMIT
ARG BIN_NAME
ARG DESCRIPTION

LABEL \
    org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.version="${BUILD_VER}" \
    org.opencontainers.image.revision="${GIT_COMMIT}" \
    org.opencontainers.image.vendor="${VENDOR}" \
    org.opencontainers.image.description="${DESCRIPTION}" \
    org.opencontainers.image.title="${BIN_NAME}" \
    com.mcg.container.help="docker exec -it <CONTAINER> /app/${BIN_NAME} --help" \
    com.mcg.container.os="scratch"

WORKDIR /app
COPY --from=build /src/bin/ /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Entry point will be added by the build.sh script
