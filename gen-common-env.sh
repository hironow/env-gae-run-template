#!/usr/bin/env bash
set -eu

case $1 in
  "dev" | "develop" | "development" )         APP_ENV=dev envsubst < ./common.dev.env > ./common.env ;;
  "stg" | "staging" )                         APP_ENV=stg envsubst < ./common.stg.env > ./common.env ;;
  "prd" | "prod" | "production" | "master" )  APP_ENV=prd envsubst < ./common.prd.env > ./common.env ;;
  * ) echo "non target... $1" && exit 1;
esac