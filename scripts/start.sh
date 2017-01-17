#!/bin/bash

script_dir="${0%/*}"
cd $script_dir/..

db_port=5432

while getopts 'p:' flag; do
  case "${flag}" in
    p) db_port="${OPTARG}" ;;
  esac
done

sed "s/{{DB_PORT}}/$db_port/g" docker-compose.tmpl > docker-compose.yml

echo "Stopping existing containers..."
docker-compose down

echo "Building and starting parrot containers..."
docker-compose up --build --force-recreate -d

echo "Attaching to logs..."
docker-compose logs -f