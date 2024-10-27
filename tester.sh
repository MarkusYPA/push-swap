#!/bin/bash

# All good test cases
cases=("" "2 1 3 6 5 8" "0 1 2 3 4 5" "0 one 2 3" "1 2 2 3" "12 26 97 81 53" "33 3 22 76 9" "example07.txt")

# Looping through them
for case in "${cases[@]}"
do
    echo
    echo "Testing: $case"
    echo
    go run . "$case"
    echo
    go run . "$case" | wc -l | xargs echo -n ; echo " instruction(s)"
    echo "-----------------------------"
done
