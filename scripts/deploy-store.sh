#!/bin/bash
source $(pwd)/scripts/cf.cfg
source $(pwd)/scripts/cf.sh

cf push ${STORE_NAME} -o markstgodard/stuff:v0 --no-start --no-manifest
cf_setenv ${STORE_NAME} stuff v1 8080

cf start ${STORE_NAME}
