#!/bin/bash

input=$1

if [ -z $input ]; then
  echo "starting..."

  docker-compose up -d --build 
  docker-compose ps -a

elif [ $input == "restart" ]; then
  echo "restarting..."

  docker-compose down 2> /dev/null
  docker-compose up -d --build 
  docker-compose ps -a

elif [ $input == "down" ]; then
  echo "stopping..."

  docker-compose down 2> /dev/null
  
fi
