#!/bin/sh


sudo echo "27" > /sys/class/gpio/export
sudo echo "22" > /sys/class/gpio/export

sudo echo "out" > /sys/class/gpio/gpio27/direction
sudo echo "out" > /sys/class/gpio/gpio22/direction


sudo chmod 777 -R /sys/class/gpio/gpio27
sudo chmod 777 -R /sys/class/gpio/gpio22
