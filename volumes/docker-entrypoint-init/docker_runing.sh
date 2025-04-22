#!/bin/bash

for i in $(seq 1 10);
do
sleep 2
if [ -n "$(docker ps --filter "name=tempy" --filter "status=running" -q)" ]; then
  break
fi
done