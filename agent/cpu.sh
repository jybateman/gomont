#!/bin/bash

CPUS=`cat /proc/cpuinfo  | grep MHz | cut -d":" -f2`

NUM=`echo $CPUS | wc -w`

OP=`echo $CPUS | tr " " "+" | xargs echo "scale=4;" | bc`

echo "scale=4; $OP/$NUM" | bc

exit 0
