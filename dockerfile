ARG GO_VERSION=1.21
ARG ALPINE_VERSION=3.19

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build

ENV CGO_ENABLED=0 GO_ENABLED=0

ARG BUILD_DATE
ARG BUILD_VER
ARG GIT_COMMIT
ARG TARGET

RUN apk --no-cache add ca-certificates

WORKDIR /src
COPY . .

RUN apk add --no-cache git && \
    go mod tidy && \
    go build -ldflags \
    "-X '{{module_name}}/cmd/conf.buildDate=${BUILD_DATE}' \
     -X '{{module_name}}/cmd/conf.buildVer=${BUILD_VER}' \
     -X '{{module_name}}/cmd/conf.buildCommit=${GIT_COMMIT}' -s -w" \
    -o ./bin/{{bin_name}} ${TARGET}

FROM scratch as final

ARG BUILD_DATE
ARG BUILD_VER
ARG GIT_COMMIT

LABEL \
    org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.version="${BUILD_VER}" \
    org.opencontainers.image.revision="${GIT_COMMIT}" \
    org.opencontainers.image.vendor="{{vendor_name}}" \
    org.opencontainers.image.description="{{description}}" \
    org.opencontainers.image.title="{{bin_name}}" \
    com.mcg.container.help="docker exec -it <CONTAINER> /app/{{bin_name}} --help" \
    com.mcg.container.os="scratch"

WORKDIR /app
COPY --from=build /src/bin/ /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/{{bin_name}}"]
