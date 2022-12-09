#!/usr/bin/env zsh

set -e

if [ $# -lt 1 ]; then
    echo "Usage: ./k6-in-docker.sh <SCRIPT_NAME> [additional k6 args]"
    exit 1
fi

# By default, we're assuming you created the extended k6 image as "k6-extended:latest".
# If not, override the name on the command-line with `IMAGE_NAME=...`.
IMAGE_NAME=${IMAGE_NAME:="k6-extended:latest"}

# Define the environment file containing variables to be passed along to the k6 container. 
ENV_FILE=${ENV_FILE:="script_vars.env"}

# Network to which the docker execution should attach, e.g. `docker network ls` to list them.
NETWORK=${NETWORK:="host"}

# Each execution is provided a unique `testid` tag to differentiate discrete test runs.
# (Not required, but provided for convenience)
SCRIPT_NAME=$1
TAG_NAME="$(basename -s .js $SCRIPT_NAME)-$(date +%s)"

# This is a basic wrapper to run a clean docker container
#   --env-file : filename providing environment variables for use by k6 and extensions
#   --network  : if dependencies are in another container, this is the network they're in
#   -v         : we're allowing scripts to be located in the current directory, or any of its children
#   -it        : running interactively
#   --rm       : once the script completes, the container will be removed (good housekeeping, you'll thank me)
#
# Anything after the $IMAGE_NAME are passed along for the k6 binary.
docker run --env-file $ENV_FILE --network $NETWORK \
 -v $PWD:/scripts -it --rm $IMAGE_NAME \
 run /scripts/$SCRIPT_NAME --tag testid=$TAG_NAME ${@:2}
