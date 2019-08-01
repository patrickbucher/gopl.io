#!/bin/sh

go build xmltree.go && cat company.xml | ./xmltree
rm -f xmltree
