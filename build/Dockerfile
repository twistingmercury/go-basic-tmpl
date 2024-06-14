ARG GOLANG_IMAGE_VERSION

FROM golang:${GOLANG_IMAGE_VERSION} as build
ENV CGO_ENABLED=0 GO_ENABLED=0

ARG BUILD_DATE
ARG BUILD_VER

RUN apk --no-cache add ca-certificates

WORKDIR /token_go_bin
COPY .. .

RUN apk add --no-cache git
RUN go mod tidy
RUN go build -ldflags \
    "-X 'token_go_module/conf.buildDate=${BUILD_DATE} -X 'token_go_module/conf.buildVer=${BUILD_VER}' -s -w" \
    -o ./bin/token_bin main.go

FROM scratch as final

ARG BUILD_DATE
ARG BUILD_VER

LABEL \
    org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.version="${BUILD_VER}" \
    org.opencontainers.image.vendor="token_docker_vendor_name" \
    org.opencontainers.image.description="token_docker_descr" \
    org.opencontainers.image.title="token_bin" \
    com.mcg.container.help="docker exec -it <CONTAINER> /app/token_go_bin --help" \
    com.mcg.container.os="scratch"

WORKDIR /app
COPY --from=build /src/bin/ /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/token_bin"]
