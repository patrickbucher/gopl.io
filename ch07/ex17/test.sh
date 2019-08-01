#!/bin/sh

go build xmlselect.go && cat company.xml | ./xmlselect 'usefulness="inquestionable"' 'salary="small"'
rm -f xmlselect
