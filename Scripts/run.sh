#/bin/bash
if [ $1 = "all" ]; then
	sh ./i/*.sh
else
	COUNTER=$1
	 while [  $COUNTER -lt $2 ]; do
	     sh ./i/$COUNTER.sh
	     let COUNTER=COUNTER+1 
	 done
fi