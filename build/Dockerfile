FROM golang:alpine as build
ENV CGO_ENABLED=0 GO_ENABLED=0

ARG BUILD_DATE

RUN apk --no-cache add ca-certificates

WORKDIR /src
COPY .. .

RUN apk add --no-cache git
RUN go mod tidy
RUN go build -ldflags "-X 'my/module/internal/conf.buildDate=${BUILD_DATE}' -s -w" -o ./bin/myapp cmd/main.go

FROM scratch as final

ARG BUILD_DATE
ARG BUILD_VER

LABEL \
    org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.version="${BUILD_VER}" \
    org.opencontainers.image.vendor="token_docker_vendor_name" \
    org.opencontainers.image.description="token_docker_descr" \
    org.opencontainers.image.title="token_go_bin"

WORKDIR /app
COPY --from=build /src/bin/ /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8181 9090

ENTRYPOINT ["/app/token_go_bin"]
