FROM alpine:3.15.2

ARG MAJOR_VERSION=1.0
ARG CSAPI_VERSION=${MAJOR_VERSION}.1

LABEL org.opencontainers.image.authors="Jorge Lopez <lowlifebob@gmail.com>" \
      org.opencontainers.image.description="CSAPI build base image contains all required packages to build CSAPI images" \
      org.opencontainers.image.title="ContainerSolutionsAPI build base" \
      org.opencontainers.image.version="${CSAPI_VERSION}"


RUN apk add --no-cache ca-certificates

# ENV PATH /usr/local/go/bin:$PATH

RUN set -eux && \
    INSTALL_PKGS="bash \
            go" && \
    apk add \
            --no-cache \
            --clean-protected \
        ${INSTALL_PKGS} && \
    rm -rf /var/cache/apk/*

