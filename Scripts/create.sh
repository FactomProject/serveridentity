
#/bin/bash

COUNTER=$1
 while [  $COUNTER -lt $2 ]; do
     nohup serveridentity full Es2Rf7iM6PdsqfYCo3D1tnAR65SkLENyWJG1deUzpRMQmbh9F3eG -n="$COUNTER" &
     let COUNTER=COUNTER+1 
 done

