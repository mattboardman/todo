# todo
A To-Do Application

## Features - Summary
* Read - Gets all ToDo items
* Create - Adds a new ToDo item
* Update - Updates an existing ToDo item
* Delete - Deletes an existing ToDo item
* Search - Finds a matching ToDo item with the input 
  
## Deploy
#### Build from source

````bash
git clone https://github.com/mattboardman/todo
cd todo
go build .
./todo
````
Then navigate your browser to the default http://localhost:1080

#### Deploy as Docker image

Docker image hosted here: https://hub.docker.com/r/mattboardman/todo/

````bash
docker pull mattboardman/todo:first
docker run -it -p <frontend-port>:1080 -p <backend-port>:8080 mattboardman/todo:first
````
Then navigate your browser to http://localhost:<frontend-port\>


