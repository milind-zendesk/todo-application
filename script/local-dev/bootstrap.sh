#!/usr/bin/env bash
set -e

# This script checks each of the tools we expect to be present on the local dev environment and installs them if they're not present

# Colours!
RED=$'\x1b[1;31m'
GREEN=$'\x1b[1;32m'
CYAN=$'\x1b[1;36m'
GREY=$'\x1b[0;90m'
NORM=$'\x1b[0m'

# MySQL client
mysql_client_version=$(mysql --version 2>/dev/null | grep "Ver" | sed "s/.*Ver //")

if [[ -z "$mysql_client_version" ]]; then
  echo "🥁  ${RED}No Mysql client version detected, installing...${GREY}"
  brew install mysql-client
  brew link mysql-client --force
else
  echo -e "🥁  ${NORM}Mysql client version detected:\t${GREEN}${mysql_client_version}${NORM}"
fi