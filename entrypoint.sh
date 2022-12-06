#!/bin/bash
sleep 5
if ! ./cmd --init; then
   echo "Init redis failed. Exiting..."
else
   exec ./cmd
fi