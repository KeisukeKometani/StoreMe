#!/bin/sh

if [ $1 = 'generate' ] || [ $1 = 'g' ]; then
  shift
  migrate create -ext sql -dir /db/migrate -format "20060102150405" $@
elif [ $1 = 'migrate' ]; then
  shift
  migrate -source "file:///db/migrate" \
    -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" \
    $@
elif [ $1 = 'raw' ] || [ $1 = 'r' ]; then
  shift
  migrate $@
elif [ $1 = 'noop' ]; then
  exit 0
else
  echo "No implement command '$1'"
  exit 1
fi
