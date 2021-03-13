# TcbHomework

## Informations

### Version

0.1.0

## Content negotiation

### URI Schemes

- http

### Consumes

- application/json

### Produces

- application/json

## All endpoints

### pool

| Method | URI              | Name                        | Summary |
| ------ | ---------------- | --------------------------- | ------- |
| POST   | /api/pools/add   | [insert pool](#insert-pool) |         |
| POST   | /api/pools/query | [querry pool](#querry-pool) |         |

## Paths

### <span id="insert-pool"></span> insert pool (_insertPool_)

```
POST /api/pools/add
```

insert/append values to pools

#### Parameters

| Name | Source | Type                       | Go type             | Separator | Required | Default | Description |
| ---- | ------ | -------------------------- | ------------------- | --------- | :------: | ------- | ----------- |
| body | `body` | [PoolObject](#pool-object) | `models.PoolObject` |           |          |         |             |

#### All responses

| Code                            | Status | Description            | Has headers | Schema                                |
| ------------------------------- | ------ | ---------------------- | :---------: | ------------------------------------- |
| [200](#insert-pool-200)         | OK     | OK                     |             | [schema](#insert-pool-200-schema)     |
| [default](#insert-pool-default) |        | generic error response |             | [schema](#insert-pool-default-schema) |

#### Responses

##### <span id="insert-pool-200"></span> 200 - OK

Status: OK

###### <span id="insert-pool-200-schema"></span> Schema

[PoolObjectAddResponse](#pool-object-add-response)

##### <span id="insert-pool-default"></span> Default Response

generic error response

###### <span id="insert-pool-default-schema"></span> Schema

| Name                   | Type   | Go type | Default | Description | Example |
| ---------------------- | ------ | ------- | ------- | ----------- | ------- |
| insertPool defaultBody | string |         |         |             |         |

### <span id="querry-pool"></span> querry pool (_querryPool_)

```
POST /api/pools/query
```

insert/append values to pools

#### Parameters

| Name | Source | Type                                    | Go type                   | Separator | Required | Default | Description |
| ---- | ------ | --------------------------------------- | ------------------------- | --------- | :------: | ------- | ----------- |
| body | `body` | [PoolQueryRequest](#pool-query-request) | `models.PoolQueryRequest` |           |          |         |             |

#### All responses

| Code                            | Status | Description            | Has headers | Schema                                |
| ------------------------------- | ------ | ---------------------- | :---------: | ------------------------------------- |
| [200](#querry-pool-200)         | OK     | OK                     |             | [schema](#querry-pool-200-schema)     |
| [default](#querry-pool-default) |        | generic error response |             | [schema](#querry-pool-default-schema) |

#### Responses

##### <span id="querry-pool-200"></span> 200 - OK

Status: OK

###### <span id="querry-pool-200-schema"></span> Schema

[PoolQueryResponse](#pool-query-response)

##### <span id="querry-pool-default"></span> Default Response

generic error response

###### <span id="querry-pool-default-schema"></span> Schema

| Name                   | Type   | Go type | Default | Description | Example |
| ---------------------- | ------ | ------- | ------- | ----------- | ------- |
| querryPool defaultBody | string |         |         |             |         |

## Models

### <span id="error-response"></span> errorResponse

| Name          | Type   | Go type | Default | Description | Example |
| ------------- | ------ | ------- | ------- | ----------- | ------- |
| errorResponse | string | string  |         |             |         |

### <span id="pool-object"></span> poolObject

**Properties**

| Name       | Type                       | Go type   | Required | Default | Description | Example |
| ---------- | -------------------------- | --------- | :------: | ------- | ----------- | ------- |
| poolId     | int64 (formatted number)   | `int64`   |    ✓     |         |             |         |
| poolValues | []int32 (formatted number) | `[]int32` |    ✓     |         |             |         |

### <span id="pool-object-add-response"></span> poolObjectAddResponse

> return add status, staus fileld value will be either "appended" or "inserted"

**Properties**

| Name   | Type   | Go type  | Required | Default | Description | Example |
| ------ | ------ | -------- | :------: | ------- | ----------- | ------- |
| status | string | `string` |    ✓     |         |             |         |

### <span id="pool-query-request"></span> poolQueryRequest

**Properties**

| Name       | Type                      | Go type   | Required | Default | Description | Example |
| ---------- | ------------------------- | --------- | :------: | ------- | ----------- | ------- |
| percentile | double (formatted number) | `float64` |    ✓     |         |             |         |
| poolId     | int64 (formatted number)  | `int64`   |    ✓     |         |             |         |

### <span id="pool-query-response"></span> poolQueryResponse

**Properties**

| Name               | Type                     | Go type | Required | Default | Description | Example |
| ------------------ | ------------------------ | ------- | :------: | ------- | ----------- | ------- |
| calculatedQuantile | int32 (formatted number) | `int32` |    ✓     |         |             |         |
| totalCount         | int64 (formatted number) | `int64` |    ✓     |         |             |         |
