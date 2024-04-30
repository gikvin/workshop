# workshop
From local development to cloud delivery

> **WARNING**  
> This is a simple example of a Golang CRUD application that connects to a MySQL database. Please note that this code is for demonstration purposes only and contains bugs. It doesn't strictly follow Go style conventions and isn't intended for production environments.

## Table of Contents
- [Pre-requisites](#pre-requisites)
- [How to start](#how-to-start)
- [How to test](#how-to-test)
    - [Create](#create)
    - [Read](#read)
    - [Update](#update)
    - [Delete](#delete)
- [Known bugs](#known-bugs)

## Pre-requisites
following tools are needed to start using this repository
* Git
* Golang (tested version 1.22)
* Docker
* Docker compose
* Docker Desktop (optional, if you have docker desktop installed, it automatically installs docker compose)

## How to start
* use docker compose to start the application
```bash
docker compose up
```

## How to test
Application supports following http methods: GET, POST, PUT, DELETE
here are examples for Creating, Reading, Updating and Deleting User

### Create
```bash
curl -X POST -H "Content-Type: application/json" -d '{"Name":"Tony Stark","Email":"tony@example.com"}' http://localhost:8090/user
```

### Read
* To list all users
```bash
curl http://localhost:8090/users
```
* To list specific user with ID
```bash
curl http://localhost:8090/user/{ID}
```

### Update
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"Name":"Steve Rogers","Email":"steve@example.com"}' http://localhost:8090/user/{ID}
```

### Delete
```bash
curl -X DELETE -H "Content-Type: application/json" http://localhost:8090/user/1
```


## Known bugs
* Update and Delete method returns http 200 with successfully deleted message if ID does not exist
* many more
