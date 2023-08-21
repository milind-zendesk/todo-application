# todo-application
This is a CRUD application project.

## Introduction
This CLI application lets you play with 2 tables in MySQL with the help of different APIs. You can perform all the CRUD operations and play with the data.

## Before you start

### Clone This Repo
Open a desired location in your terminal and paste this command.
`git@github.com:milind-zendesk/todo-application.git`
Once you've cloned the repo, open the project in any editor of your choice. As this project is build on Golang, VSCode or GoLand editor should be preferred. 

### Installing Requirements
Open the terminal of your editor and hit the following command
`make bootstrap`
This command will install all the requirements that are needed to run this project and now you are ready to go!

## Run the Project
Once all your requirements are installed and the project setup is done, you can run the project using `make` or `make run` command. This command will run your application as well as create the tables in the Docker Container.
The API server will start running on port 8080. You can visit this port using localhost like this
`localhost:8080/{api}`

## APIs you can use in this application

### POST /bulk_load
This API will load the Fake/Random data in your database.

### GET /users
This API will fetch all the Users data from the `user` table.

### GET /users/{id}
This API will fetch the data of a particular User.

### POST /add_user
This API will add the user data in the Database. You need to pass the payload in this format.
`'{"name":"Milind Shinde", "location":"Pune"}'`

### PUT /edit_user/{id}
This API lets you edit the existing user. This will require the payload in below format.
`'{"name":"Milind Shinde", "location":"Pune"}'`

### GET /todos
This API will fetch all the data from the `todo` table.

### GET /todos/{id}
This API will fetch the data of a particular Todo.

### POST /insert_todo
This API will add the todo data in the todo table. You need to pass the payload in this format.
`'{"title":"Shopping", "status":"done", "priority":"high", "user_id": 2}'` 
Here, user_id is the ID from the user table.

### PUT /edit_todo/{id}
This API lets you edit the existing todo data. This will require the payload in below format.
`'{"title":"Shopping", "status":"done", "priority":"high", "user_id": 2}'` 

### DELETE /delete_todo/{id}
This API lets you delete any record from the table.

### GET /user_todos/{id}
This API will fetch the record of a particular user with all his Todos with the statistics of their priorities. you need to pass the user_id with this API.

## Running the Tests
For running the test cases for all the APIs, you can use the following command
`make test`