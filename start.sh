#!/bin/bash

# turn on bash's job control
set -m

# Start the primary process and put it in the background
./cmd/main &

# Start the helper process
./data_simulator/main

# the my_helper_process might need to know how to wait on the
# primary process to start before it does its work and returns


# now we bring the primary process back into the foreground
# and leave it there
fg %1
############################
# Start the first process
#/cmd & # your first application
#P1=$!
#/cmd2 & # your second application
#P2=$!
#wait $P1 $P2