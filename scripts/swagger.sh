#!/usr/bin/env bash

POST_BODY=$(swagger generate spec)

SWAGGER_V3=$(curl --silent --location --request POST \
  https://converter.swagger.io/api/convert \
  --header 'Content-Type: application/json' \
  --header 'Accept: application/yaml' \
  --data "$POST_BODY")

TIMES_FROM="2021-10-19T13:42:35+02:00"

TIMES_TO="\"2021-10-19T13:42:35+02:00\""

SWAGGER_V3=$(python ./scripts/replace.py "$SWAGGER_V3" "$TIMES_FROM" "$TIMES_TO")

echo "$SWAGGER_V3" > ./api/swagger.yaml
