#!/bin/bash

rounds=$1

host=$(oc get route go-app-route --no-headers | awk '{print $2}')
curl -k -c /tmp/cookie.txt http://${host}/path

echo "Running $rounds executions on $host with cookie:"
cat /tmp/cookie.txt
echo "########"
for i in $(seq 1 ${rounds})
do
    response=$(curl -s -k -b /tmp/cookie.txt http://${host}/path)
    echo ${response}
done
podName=$(echo $response | cut -d' ' -f4)
echo "Killing Pod ${podName}"
oc delete pod ${podName}

for i in $(seq 1 ${rounds})
do
    curl -k -b /tmp/cookie.txt http://${host}/path
done
