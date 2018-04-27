#!/bin/bash

curl -X PUT -H "Content-Type: application/json" -d '{"id":"44f52c01-ddf2-459d-be19-44c057719f74", "title":"New Title"}' http://localhost:8080/v1/todo
