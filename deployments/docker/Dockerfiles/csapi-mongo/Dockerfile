ARG MAJOR_VERSION=1.0
ARG CSAPI_VERSION=${MAJOR_VERSION}.1
ARG BUILD_BASE_IMAGE=csapi-build-csapi:alpine-${CSAPI_VERSION}

FROM ${BUILD_BASE_IMAGE} as builder

FROM alpine:3.15.2

ARG MAJOR_VERSION
ARG CSAPI_VERSION


LABEL org.opencontainers.image.authors="Jorge Lopez <lowlifebob@gmail.com>" \
      org.opencontainers.image.description="CSAPI API restfull based on golang with mongoDB database" \
      org.opencontainers.image.title="ContainerSolutionsAPI API restfull" \
      org.opencontainers.image.version="${CSAPI_VERSION}"

STOPSIGNAL SIGTERM


COPY --from=builder ["/home/myapp/csapi-${CSAPI_VERSION}", "/usr/local/bin"]

RUN set -eux && \
    addgroup \
            --system \
            --gid 1995 \
        mygroup && \
    adduser \
            --system \
            --gecos "API restfull CRUD" \
            --disabled-password \
            --uid 1997 \
            --ingroup mygroup \
            --shell /sbin/nologin \
            --home /home/myuser/ \
        myuser && \
    adduser myuser root

EXPOSE 3333/TCP

WORKDIR /home/myuser

COPY ["./deployments/docker/Dockerfiles/csapi-mongo/csapi-entrypoint.sh", "/usr/bin/"]

USER 1997

ENTRYPOINT ["/usr/bin/csapi-entrypoint.sh"]
