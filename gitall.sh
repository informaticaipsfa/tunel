#!/bin/sh
COMANDO="log"

cd $1
if [ $2 != $COMANDO ]; then
	echo $(git pull origin master);
else
	git log;
fi;
