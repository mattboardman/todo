#!/bin/bash

curl -X PUT -H "Content-Type: application/json" -d '{"id":"00000000-0000-0000-0000-000000000000", "title":"New Title"}' http://localhost:8080/v1/todo
