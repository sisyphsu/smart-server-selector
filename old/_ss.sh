#!/usr/bin/env bash

servers=()
servers+=( test, 112.111.111.111, grpc )
servers+=( test, 112.111.111.112, sock )

export IFS=","

# parse server string to map
function parseServer() {
    declare -a arr
    for word in ${1}; do
        arr+=(${word})
    done
    echo ${arr[*]}
}

for row in ${servers[@]};
do
    parseServer ${row}
done
