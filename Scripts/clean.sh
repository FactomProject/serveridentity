#/bin/bash

COUNTER=$1
rm nohup.out
 while [  $COUNTER -lt $2 ]; do
     rm "./i/$COUNTER.sh"
     let COUNTER=COUNTER+1 
 done
