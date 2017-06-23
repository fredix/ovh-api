#!/bin/bash

TYPE=$1
NAME=$2
STATE=$3

case $STATE in
        "MASTER") /usr/local/bin/move_ip -ip IP_FAILOVER
                  exit 0
                  ;;
        *)        echo "unknown state"
                  exit 1
                  ;;
esac
