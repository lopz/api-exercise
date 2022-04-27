FROM alpine:3.15.2

ARG MAJOR_VERSION=1.0
ARG CSAPI_VERSION=${MAJOR_VERSION}.1

LABEL org.opencontainers.image.authors="Jorge Lopez <lowlifebob@gmail.com>" \
      org.opencontainers.image.description="CSAPI build base image for populate Database" \
      org.opencontainers.image.title="ContainerSolutionsAPI populate DB" \
      org.opencontainers.image.version="${CSAPI_VERSION}"


#RUN apk add --no-cache ca-certificates

RUN set -eux && \
    INSTALL_PKGS="bash \
            ca-certificates \
            mongodb-tools \
            py3-pip \
            python3" && \
    apk add \
            --no-cache \
            --clean-protected \
        ${INSTALL_PKGS} && \
    rm -rf /var/cache/apk/*

WORKDIR /home/myapp

COPY ["./data/mongo/titanic.csv", "/tmp/"]
COPY ["./deployments/docker/Dockerfiles/mongo-seed/seeder.py", "/tmp/"]
COPY ["./deployments/docker/Dockerfiles/mongo-seed/seeder-entrypoint.sh", "/usr/bin/"]

ENTRYPOINT ["/usr/bin/seeder-entrypoint.sh"]