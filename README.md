# TCB Homework

## Prerequisite

In order to run/test/build the solution of this homework, please install Golang using instructions in below link:
https://golang.org/doc/install

To makesure that you have GoLang installed on your PC, you should be able to execute below command :

```
go version
```

sample output:

```
go version go1.15.2 darwin/amd64
```

- Internet access should be available. Some go modules will be automatically downloaded if not exist in your PC when run/build this solution

## Description:

This repo delivers the solution of creating a Rest API application with two post endpoints. One is for inserting and the other is for querying pool object.
The pool object consists of an id and an array containing pool values, where it could be inserted or appended.
The query return calculated quantile of a given poolid and percentile.

- For endpoints detail description: Please read **_ENDPOINTS-DOC.md_**
- The solution is written in Golang
- The solution is based on go-swagger framework (an implementation of Open-API 2.0 on Golang).

  - The endpoints specification are descrided in a file called **_swagger.yml_** using Open-API 2.0 standard. Then boilerplate codes and documentation are generated from this **_swagger.yml_** file.

  - After that the detail implementation are applied based on the generated boilerplate codes.
  - More about go-swagger framework at https://goswagger.io/

- Pool objects are stored in a Map[poolId]PoolObject. The Insert and query are thread safe using sync RW mutex. Data will be lost when server shutting down due to no database implementation

## Project structure

```
.
├── README.md
├── ENDPOINTS-DOC.md (generated Doc from swagger.yml)
├── apihandler (detail implementation)
├── cmd
│   └── poolservice-server
│       └── main.go (entry point of the REST API application)
├── genAPI.sh (generate boilate codes from swagger.yml file command; must install swagger to use this )
├── go.mod
├── go.sum
├── models (boilerplate code from swagger)
├── restapi ( boilerplate code from swagger)
├── storage ( pool objects storage and query functions)
├── swagger.yml
└── util (helper functions such as quantile calculation)
```

## Run the homework solution

navigate to < PATH TO REPO >/cmd/poolservice-server folder, then run:

```
go run main.go --port 5000
```

The command will spin up a REST API server on localhost port 5000 to serve the endpoints
if port is not specified, a random port will be assigned

sample output:

```
2021/03/13 13:46:21 Serving poolservice at http://127.0.0.1:5000
```

## Build executable binary file

Navigate to < PATH TO REPO >/cmd/poolservice-server folder, then run:

```
go build
```

a new "poolservice-server" binary file is generated in the same folder.
Give this file executable permission and run it directly with below command to serve the endpointds, for example on port 5000

```
./poolservice-server --port 5000
2021/03/13 14:28:19 Serving poolservice at http://127.0.0.1:5000
```

## Unit Tests

navigate to root folder of this repo, from there run below command:

```
go test ./util
go test ./storage
go test ./apihandlertest

```

- The unit tests cover our written codes, the generated codes that are in "restapi", "models" folders are not tested
- tests are written in \*\_test.go files

- For more verbose output, please add -v option as below command:

```
go test ./util -v
go test ./storage -v
go test ./apihandlertest -v

```
