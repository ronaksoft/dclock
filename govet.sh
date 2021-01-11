#!/usr/bin/env bash

# Generate codes
go generate ./... || exit

# Make sure the code guide lines are met
go vet ./... || exit

# Format the code
go fmt ./... || exit