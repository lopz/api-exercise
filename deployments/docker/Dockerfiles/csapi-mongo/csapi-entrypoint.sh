#!/bin/sh

set -o pipefail

set +e

MAJOR_VERSION="1.0"
CSAPI_VERSION="$MAJOR_VERSION.1"

echo "** Starting CS API..."
exec /usr/local/bin/csapi-${CSAPI_VERSION}