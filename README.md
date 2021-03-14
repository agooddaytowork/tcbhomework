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

## Description

This repo delivers the solution of creating a Rest API application with two post endpoints. One is for inserting and the other is for querying pool object.

- The pool object consists of an id and an array containing pool values.
  - The pool object values can be inserted, then later appended
  - The query return calculated quantile, total value counts of a given poolid and percentile.
- input validation is implemented. It returns related http status codes with a string message for more details.
- For endpoints detail description: Please read **_ENDPOINTS-DOC.md_**
- The solution is based on go-swagger framework (an implementation of Open-API 2.0 on Golang).

  - The endpoints specification are descrided in a file called **_swagger.yml_** using Open-API 2.0 standard. Then boilerplate codes and documentation are generated from this **_swagger.yml_** file.

  - After that the detail implementation are applied based on the generated boilerplate codes.
  - More about go-swagger framework at https://goswagger.io/

- Pool objects are stored in a Map[poolId]PoolObject. The Insert and query are thread safe using sync RW mutex.
- Data will be lost when server shutting down due to no database implementation

## Project structure

```
.
├── README.md
├── ENDPOINTS-DOC.md (generated Doc from swagger.yml)
├── apihandler (detail implementation)
|── apihandlertest (dedicated test package for apihander)
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

Navigate to root folder of this repo, from there run below command:

```
go run cmd/poolservice-server/main.go --port 5000
```

The command will spin up a REST API server on localhost, port 5000, to serve the endpoints.

If port is not specified, a random port will be assigned

sample output:

```
2021/03/13 13:46:21 Serving poolservice at http://127.0.0.1:5000
```

### Example Usages

- Assume that the server is running at http://127.0.0.1:5000

- add a pool object

```
curl -d '{"poolId":12345,"poolValues":[1,2,3,4,5]}`' -H "Content-Type: application/json" -X POST http://127.0.0.1:5000/api/pools/add

> {"status":"inserted"}

curl -d '{"poolId":12345,"poolValues":[1,2,3,4,5]}`' -H "Content-Type: application/json" -X POST http://127.0.0.1:5000/api/pools/add

> {"status":"appended"}

curl -d '{"poolId":12345,"poolValues":[1,2,3,4,5]}`' -H "Content-Type: application/json" -X POST http://127.0.0.1:5000/api/pools/add

> {"status":"appended"}

curl -d '{"polId":12345,"poolVaes":[2]}`' -H "Content-Type: application/json" -X POST http://127.0.0.1:5000/api/pools/add

> poolId in body is required
```

- query a pool object

```
curl -d '{"poolId":12345,"percentile":95.5}`' -H "Content-Type: application/json" -X POST http://127.0.0.1:5000/api/pools/query
> {"calculatedQuantile":5,"totalCount":15}

curl -d '{"poolId":2222,"percentile":95.5}`' -H "Content-Type: application/json" -X POST http://127.0.0.1:5000/api/pools/query
> "Pool with id 2222 not found"

curl -d '{"poId":2222,"perceile":95.5}`' -H "Content-Type: application/json" -X POST http://127.0.0.1:5000/api/pools/query
> percentile in body is required
```

## Build executable binary file for deployment

Navigate to < PATH TO REPO >/cmd/poolservice-server folder, then run:

```
go build
```

a new "poolservice-server" binary file is generated in the same folder.
Give this file executable permission and run it directly with below command to serve the endpointds, for example on port 5000

```
chmod +x poolservice-server
./poolservice-server --port 5000
> 2021/03/13 14:28:19 Serving poolservice at http://127.0.0.1:5000
```

## Unit Tests

navigate to root folder of this repo, from there run below command:

```
go test ./util
go test ./storage
go test ./apihandlertest

```

sample output:

```
go test ./util
go test ./storage
go test ./apihandlertest
ok      tam.io/homework/util    (cached)
ok      tam.io/homework/storage (cached)
ok      tam.io/homework/apihandlertest  0.033s
```

- The unit tests cover only our written codes, the generated codes that are in "restapi", "models" folders are not tested
- tests are written in \*\_test.go files
- Regarding apihandlertest, a httptest server is spin up to make the test
- For more verbose output, please add -v option as below command:

```
go test ./util -v
go test ./storage -v
go test ./apihandlertest -v

```

## Further developments

- Add caching for calculated result. At the moment the query API need to calculate on every request
- Sort the input pool values right after inserted/appended. At the moment, the sort happens inside quantile calculation functions.
- Add benchmark, refactor if required
- Add Flag to configure the maximum length of input poolValues. It is fixed to 1M points at the moment.
- Add util scripts for build, test and deploys
