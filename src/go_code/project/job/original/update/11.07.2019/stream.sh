#!/bin/bash

while IFS='' read -r line || [[ -n "$line" ]]; do
    foo=$line
    for (( i=0; i<${#foo}; i++ )); do
        echo -n "${foo:$i:1}"
        sleep 0.04
    done
    echo
    sleep 10
done
