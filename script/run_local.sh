#!/bin/bash

TARGET_BUILD_PLATFORM=linux/amd64

./script/build.sh ${TARGET_BUILD_PLATFORM} && docker-compose up -d