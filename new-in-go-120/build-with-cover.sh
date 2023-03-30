#!/usr/bin/env bash

# START OMIT
go build -cover -o boundary ./cmd/boundary
mkdir coverage
GOCOVERDIR=$(pwd)/coverage ./boundary dev
go tool covdata textfmt -i=$(pwd)/coverage -o $(pwd)/profile.txt
go tool cover -html=profile.txt
# END OMIT
