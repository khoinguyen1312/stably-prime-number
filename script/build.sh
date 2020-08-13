#!/bin/bash

DOCKER_BUILDKIT=1

PLATFORM=$1
TARGET_IMAGE_NAME=$2

options=("darwin/amd64" "windows/amd64" "linux/amd64" "linux/arm" "Quit")

if [ ! $PLATFORM ] 
    then
    PS3='Please enter Platform you want to build: '
    select opt in "${options[@]}"
    do
        PLATFORM=$opt
        break
    done
fi

if [ ! $TARGET_IMAGE_NAME ] 
then
    TARGET_IMAGE_NAME='stably/prime-number:latest'
fi

echo """
==========================================================================================
 
   Building using platform '${PLATFORM}' to image '${TARGET_IMAGE_NAME}'

==========================================================================================
"""

make PLATFORM=${PLATFORM} TARGET_IMAGE_NAME=${TARGET_IMAGE_NAME}