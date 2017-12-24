#!/usr/bin/env bash

strategy=$1
n=$2

cmd="2048 ai --strategy ${strategy} --silent"
echo "Running '${cmd}' ${n} times and putting results into es index '${strategy}'."

for ((i = 0; i < $n; i++)); do
    echo ${i}
    2048 ai --strategy ${strategy} --silent | curl -XPOST "http://localhost:9200/${strategy}/2048ai_results/?pretty" -H "Content-Type: application/json" -d @-
done
