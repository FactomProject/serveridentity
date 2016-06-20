#/bin/bash

COUNTER=$1
rm nohup.out
if [ $1 = "all" ]; then
	for each in ./i/*.sh; do rm $each; done
else
 while [ $COUNTER -lt $2 ]; do
     rm "./i/$COUNTER.sh"
     let COUNTER=COUNTER+1 
 done
fi