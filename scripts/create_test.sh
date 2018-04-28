#!/bin/bash

curl -d "Content-Type: application/json" -d '{"title":"Testing", "description":"Out Posting"}' http://localhost:8080/v1/todo
