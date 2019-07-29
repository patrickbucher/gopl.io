#!/bin/sh

trap "kill 0" EXIT

go run shop.go &
sleep 1

host=localhost:8000

echo 'list initial items'
curl "${host}/list"

echo 'add a couple of items'
curl -X POST "${host}/create" -F item='belt' -F price='29.59'
curl -X POST "${host}/create" -F item='shirt' -F price='59.50'
curl -X POST "${host}/create" -F item='hat' -F price='89.99'

echo 'read some items'
curl "${host}/read?item=belt"
curl "${host}/read?item=shirt"
curl "${host}/read?item=hat"

echo 'update some items'
curl -X PATCH "${host}/update?item=belt&price=19.90"
curl -X PATCH "${host}/update?item=shirt&price=49.90"
curl -X PATCH "${host}/update?item=hat&price=79.90"
curl "${host}/read?item=belt"
curl "${host}/read?item=shirt"
curl "${host}/read?item=hat"

echo 'delete some items'
curl -X DELETE "${host}/delete?item=belt"
curl -X DELETE "${host}/delete?item=shirt"
curl -X DELETE "${host}/delete?item=hat"

echo 'list final items'
curl "${host}/list"
