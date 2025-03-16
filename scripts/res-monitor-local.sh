#!/usr/bin/env bash

n=181

top -b -d 1 -n "$n" | grep geth | awk '{print $9}' > exp/cpu-geth.txt &
top -b -d 1 -n "$n" | grep geth | awk '{print $10}' > exp/mem-geth.txt &
top -b -d 1 -n "$n" | grep trust | awk '{print $9}' > exp/cpu-tc.txt &
top -b -d 1 -n "$n" | grep trust | awk '{print $10}' > exp/mem-tc.txt &
