#!/usr/bin/env bash

curl -X GET -H "Content-Type: application/json" \
  localhost:8080/employeerisk/${1}
