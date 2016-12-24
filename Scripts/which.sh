#/bin/bash

for each in ./i/*.sh; do 
	r=$(grep -c "$1" $each)
	if [ $r -gt 1 ]; then
		echo $each
	fi
done