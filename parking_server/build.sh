#!/usr/bin/env bash

set -e

podman build \
  -f Containerfile \
  -t aren55555/grpc-example \
  --platform linux/amd64
