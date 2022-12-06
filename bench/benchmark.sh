#!/bin/bash
sleep 10
ab -c 20 -t 10 http://${benchAddr}/json/hackers