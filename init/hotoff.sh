#!/bin/sh

#port 22 is the hot relay
#port 27 is the cold relay 
arg="$1"

if [ "$arg" = "cold" ] ; then
    echo 0 > /sys/class/gpio/gpio27/value
    echo "relay off cold"
elif [ "$arg" = "hot" ] ; then
    echo 0 > /sys/class/gpio/gpio22/value
    echo "relay off hot" 
else
    echo “Wrong argument value.Use values hot or cold.”
    echo $arg
fi
