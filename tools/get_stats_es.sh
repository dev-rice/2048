#!/usr/bin/env bash

es_index=$1

curl "localhost:9200/${es_index}/_search?pretty" -H "Content-Type: application/json" -d \
'{
    "size": 0,
    "aggs" : {
        "moves_made_stats" : { "extended_stats" : { "field" : "MovesMade" } },
        "score_stats" : { "extended_stats" : { "field" : "Score" } },
        "elapsed_time_seconds_stats" : { "extended_stats" : { "field" : "ElapsedTimeSeconds" } },
        "biggest_tile_stats" : { "extended_stats" : { "field" : "BiggestTile" } }
     }
}'
