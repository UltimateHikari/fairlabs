#!/bin/bash

PREF="Fairlabs-bootstrap:"
NAME=fairlabs-server-test

echo "$PREF building.."
docker build ./server --tag server:0.1
echo "$PREF running.."
docker run --name $NAME -p 8080:8080 -d server:0.1
echo "$PREF do not forget to stop & rm container $NAME"

