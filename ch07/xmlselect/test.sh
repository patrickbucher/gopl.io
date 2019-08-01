#!/bin/sh

go build xmlselect.go && cat spec.xml | ./xmlselect div div h2
rm -f xmlselect
