#!/bin/bash

USER_ID_PREFIX="{UID $UID}"

logger $USER_ID_PREFIX started shell

exit_handler() {
	clean_up
	echo "Exited"
	logger $USER_ID_PREFIX exited
	trap - SIGHUP SIGINT SIGTERM
	kill -- -$$
}

clean_up() {
	rm -rf input.txt
}

count_tokens() {
	printf "Count[%d]: " $1 
	cat input.txt | tr '\n' " " | tr -s " " | tr " " '\n' | tr 'A-Z' 'a-z' | sort | uniq -c | sort -nr | awk '{print $2 ": " $1}' | tr "\n" " "
	echo ""
}

trap exit_handler SIGHUP SIGINT SIGTERM

if [ -z $1 ] 
	then
	num_chars=3
else
	num_chars=$1
fi
n=$((num_chars-1))
> input.txt
for ((c=0;;c++))
do
	read -p "Input:[$c] " line
	logger $USER_ID_PREFIX entered input
	echo $line | sed -e 's/[^a-zA-Z ]//g' | sed -E "s/\b\w{0,$n}\b//g" >> input.txt
	count_tokens $c
done
