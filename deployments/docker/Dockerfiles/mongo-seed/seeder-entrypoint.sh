#!/bin/sh

set -o pipefail

set +e

if [ "$1" = 'bash' ]; then
    exec "$@"
fi

exec /usr/bin/python3 /tmp/seeder.py $1
