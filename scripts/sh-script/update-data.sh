#!/usr/bin/env bash

curl -X PUT -H "Content-Type: application/json" \
  -d "@json-data/update-employee.json" \
  localhost:8080/employees
