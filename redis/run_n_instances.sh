#!/bin/bash
echo Running $1 instances
for (( i=1; i<=$1; i++ ))
do 
     echo "Instance $i"
     go run main.go &
done