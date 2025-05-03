#!/bin/bash

set -a
source .env.local
set +a

go run ./cmd/daily-init
