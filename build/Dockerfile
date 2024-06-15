FROM golang:alpine as build
ENV CGO_ENABLED=0 GO_ENABLED=0

ARG BUILD_DATE

RUN apk --no-cache add ca-certificates

WORKDIR /token_go_bin
COPY .. .

RUN apk add --no-cache git
RUN go mod tidy
RUN go build -ldflags \
    "-X 'token_go_module/internal/conf.buildDate=${BUILD_DATE}' -s -w" \
    -o ./bin/token_bin cmd/main.go

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
COPY --from=build /token_go_bin/bin/ /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/token_bin"]
