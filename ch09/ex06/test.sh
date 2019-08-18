#!/bin/sh

echo '1 Process'
GOMAXPROCS=1 go test -bench=.

echo '2 Processes'
GOMAXPROCS=2 go test -bench=.

echo '4 Processes'
GOMAXPROCS=4 go test -bench=.

echo '8 Processes'
GOMAXPROCS=8 go test -bench=.
