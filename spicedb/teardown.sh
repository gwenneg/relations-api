#!/bin/bash
# Function to check if a command is available
source ./spicedb/check_docker_podman.sh

set -e

${DOCKER} compose --env-file ./spicedb/.env -f ./docker-compose.yaml down


