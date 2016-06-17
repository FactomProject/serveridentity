#/bin/bash

COUNTER=$1
 while [  $COUNTER -lt $2 ]; do
     sh $COUNTER.sh
     let COUNTER=COUNTER+1 
 done