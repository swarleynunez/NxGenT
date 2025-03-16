#!/usr/bin/env bash

n=181

top -b -d 1 -n "$n" | grep "$1" | awk '{print $9}' > cpu-geth.txt &
top -b -d 1 -n "$n" | grep "$1" | awk '{print $10}' > mem-geth.txt &
top -b -d 1 -n "$n" | grep "$2" | awk '{print $9}' > cpu-tc.txt &
top -b -d 1 -n "$n" | grep "$2" | awk '{print $10}' > mem-tc.txt &
