#!/usr/bin/env bash

curl -X POST -H "Content-Type: application/json" \
  -d "@json-data/employee.json" \
  localhost:8080/employees

curl -X POST -H "Content-Type: application/json" \
  -d "@json-data/role.json" \
  localhost:8080/roles

curl -X POST -H "Content-Type: application/json" \
  -d "@json-data/application.json" \
  localhost:8080/applications

curl -X POST -H "Content-Type: application/json" \
  -d "@json-data/db_access.json" \
  localhost:8080/dbaccesses
