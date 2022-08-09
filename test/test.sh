#!/bin/bash

rounds=$1
host=$2
port=$3

printf "%s\tFALSE\tFALSE\t0\tpod-selector\tmyvalue\n" $host  > /tmp/cookie.txt

echo "Running $rounds executions on $host:$port with cookie:"
cat /tmp/cookie.txt
echo "########"
for i in $(seq 1 ${rounds})
do
    curl -k -b /tmp/cookie.txt http://${host}:${port}/path
done
