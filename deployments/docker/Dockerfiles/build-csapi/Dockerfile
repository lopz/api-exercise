ARG MAJOR_VERSION=1.0
ARG CSAPI_VERSION=${MAJOR_VERSION}.1
ARG BUILD_BASE_IMAGE=csapi-build-base:alpine-${CSAPI_VERSION}

FROM ${BUILD_BASE_IMAGE} as builder

ARG MAJOR_VERSION
ARG CSAPI_VERSION

LABEL org.opencontainers.image.authors="Jorge Lopez <lowlifebob@gmail.com>" \
      org.opencontainers.image.description="CSAPI build base image contains all required packages to run CSAPI" \
      org.opencontainers.image.title="ContainerSolutionsAPI build base API" \
      org.opencontainers.image.version="${CSAPI_VERSION}"


# # These commands copy your files into the specified directory in the image
# # and set that as the working location
COPY ./internal /home/myapp/internal
COPY ./go.mod ./go.sum /home/myapp/
COPY ./cmd /home/myapp/cmd

WORKDIR /home/myapp

# # This command compiles your app using GO, adjust for your source code
RUN go build -o /home/myapp/csapi-${CSAPI_VERSION} /home/myapp/cmd/main.go


