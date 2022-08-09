#!/bin/bash

rounds=$1

POD_NAME=$(oc get pod -l app=go-app -o name | head -1 | cut -d'/' -f2)
echo "Running on Pod ${POD_NAME}"
oc cp test.sh ${POD_NAME}:/tmp/test.sh
oc exec -it ${POD_NAME} -- /tmp/test.sh $1 go-app-svc 80
