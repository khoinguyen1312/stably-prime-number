#!/bin/bash

APP_NAME=stably-prime-number
PROCESS_NAME=web

TARGET_BUILD_PLATFORM=linux/amd64
TARGET_LOCAL_IMAGE_NAME=stably/prime-number:latest

./script/build.sh ${TARGET_BUILD_PLATFORM} ${TARGET_LOCAL_IMAGE_NAME} &&
echo """
==========================================================================================
 
   Deploying to Heroku with App name: '${APP_NAME}' and Process Name: ${PROCESS_NAME}

==========================================================================================
""" &&
docker tag ${TARGET_LOCAL_IMAGE_NAME} registry.heroku.com/${APP_NAME}/${PROCESS_NAME} &&
docker push registry.heroku.com/${APP_NAME}/${PROCESS_NAME}

heroku container:release ${PROCESS_NAME} -a ${APP_NAME}