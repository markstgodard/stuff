#!/bin/bash
source $(pwd)/scripts/cf.cfg

cf push ${STORE_NAME} -o markstgodard/stuff:v0 --no-start --no-manifest
cf set-env ${STORE_NAME} A8_SERVICE "stuff:v0"
cf set-env ${STORE_NAME} A8_ENDPOINT_PORT "${port}"
cf set-env ${STORE_NAME} A8_ENDPOINT_TYPE "http"
cf set-env ${STORE_NAME} A8_PROXY "true"
cf set-env ${STORE_NAME} A8_REGISTER "true"
cf set-env ${STORE_NAME} A8_REGISTRY_URL "http://${REGISTRY_NAME}.${ROUTES_DOMAIN}"
cf set-env ${STORE_NAME} A8_CONTROLLER_URL "http://${CONTROLLER_NAME}.${ROUTES_DOMAIN}"

cf start ${STORE_NAME}
