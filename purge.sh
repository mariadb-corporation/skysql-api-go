#! /bin/bash

set -eux

for file in $(cat .openapi-generator/FILES); do
    rm -f ${file}
done

rm -rf api docs
