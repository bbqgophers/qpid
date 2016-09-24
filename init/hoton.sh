#!/bin/sh

#port 22 is the hot relay
#port 27 is the cold relay 
arg="$1"

if [ "$arg" = "cold" ] ; then
    echo 1 > /sys/class/gpio/gpio27/value
    echo "relay on cold"
elif [ "$arg" = "hot" ] ; then
    echo 1 > /sys/class/gpio/gpio22/value
    echo "relay on cold"
else
    echo “Wrong argument value.Use values hot or cold.”

fi
