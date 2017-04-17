#!/bin/bash

# Replace with the name of your executable
EXEC="redis-simple-service"

if [ "$1" = "start" ]
then 
    ./$EXEC &
fi

if [ "$1" = "stop" ]
then 
    PID=$(ps -ef | grep $EXEC | grep -v grep | awk '{print $2}')

    if [[ "$OSTYPE" == "linux-gnu" ]]; then
        kill -SIGTERM $PID
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        kill $PID
    fi
fi