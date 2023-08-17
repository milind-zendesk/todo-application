#!/usr/bin/env bash

set -e
# Colours!
RED=$'\x1b[1;31m'
GREEN=$'\x1b[1;32m'
CYAN=$'\x1b[1;36m'
GREY=$'\x1b[0;90m'
NORM=$'\x1b[0m'

mysql_root_password="TempPassword"
sql_version="8.0"

current_directory="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Use settings from _environment.sh
source "${current_directory}/_environment.sh"

# Mysql actually starts twice when it starts up, once on port 0 as a temporary system to initialize things
# and once on port 3306 when it's actually going to do stuff.
# We let the logs tell us when it's ready.
wait_for_mysql() {
  while ! docker logs --since 5s "${MYSQL_CONTAINER}"  2>&1 | grep "mysql.*ready for connection.*port: 3306" ; do : ; done
}

# Check if the container exists and is running
mysql_status=$(docker ps -a --format '{{.Status}}' --filter "name=${MYSQL_CONTAINER}")

if [[ -z "$mysql_status" || $mysql_status == *"Exited"* ]] ; then
    echo "🥁  ${GREEN}Creating new MySQL container...${GREY}"

    docker run \
        --name "${MYSQL_CONTAINER}" \
        --publish "${MYSQL_PORT}":3306 \
        -e MYSQL_ROOT_PASSWORD="${mysql_root_password}" \
        --detach \
        mysql:$sql_version
    wait_for_mysql
  elif
    [[ "$mysql_status" == *"Up"*  ]] ; then
    echo "🥁  ${CYAN}MySQL is already running${GREY}"
  else
    echo "🥁  ${GREEN}Starting MySQL...${GREY}"
    docker restart "${MYSQL_CONTAINER}"
    wait_for_mysql
fi

echo "🥁  ${GREEN}Ensuring MySQL user and database exists...${GREY}"
mysql --host="${MYSQL_HOSTNAME}" --user=root --password="${mysql_root_password}" --port="${MYSQL_PORT}" << EOT
  -- Aurora setup
  CREATE DATABASE IF NOT EXISTS $AURORA_DATABASE_NAME;
  DROP USER IF EXISTS $AURORA_USERNAME;
  CREATE USER $AURORA_USERNAME IDENTIFIED BY '$AURORA_PASSWORD';
  GRANT ALL PRIVILEGES ON $AURORA_DATABASE_NAME.* TO $AURORA_USERNAME;

  CREATE TABLE IF NOT EXISTS $AURORA_DATABASE_NAME.user( id int(9) not null auto_increment, name varchar(255) not null, location varchar(255), primary key (id))ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
  CREATE TABLE IF NOT EXISTS $AURORA_DATABASE_NAME.todo( id int(9) not null auto_increment, title varchar(255) not null, status varchar(255) not null, priority varchar(255) not null, user_id int (9) not null, primary key (id), foreign key (user_id) references user(id))ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

EOT
echo "🥁  ${GREEN}MySQL settings done...${GREY}"

