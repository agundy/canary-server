Canary-Server
=============

Canary server is the Go backend for Canary, an ambient monitoring solution.

Getting Started
=============== 

Requires Go and Postgresql installed.

Once go is installed you can run

Run the following commands in a PostgreSQL shell

```
CREATE USER gorm WITH PASSWoRD 'mypassword';
CREATE DATABASE canary.db;

```

To build the Go application run:
```
go get ./...
go build
./canary-server
```


