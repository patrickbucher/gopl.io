#!/bin/sh

host=localhost:8000

echo 'list initial items'
curl "${host}/list"

echo 'add a couple of items'
curl -X POST "${host}/create" -F item='belt' -F price='13.59'
curl -X POST "${host}/create" -F item='shirt' -F price='59.50'
curl -X POST "${host}/create" -F item='hat' -F price='89.99'

echo 'read some items'
curl "${host}/read?item=belt"
curl "${host}/read?item=shirt"
curl "${host}/read?item=hat"
