// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Percentile",
    "version": "0.1.0"
  },
  "paths": {
    "/api/pools/add": {
      "post": {
        "description": "insert/append values to pools",
        "tags": [
          "pool"
        ],
        "operationId": "insertPool",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/poolObject"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/poolObjectAddResponse"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/errorResponse"
            }
          }
        }
      }
    },
    "/api/pools/query": {
      "post": {
        "description": "insert/append values to pools",
        "tags": [
          "pool"
        ],
        "operationId": "querryPool",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/poolQueryRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/poolObjectAddResponse"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/errorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "errorResponse": {
      "type": "string"
    },
    "poolObject": {
      "type": "object",
      "required": [
        "poolId",
        "poolValues"
      ],
      "properties": {
        "poolId": {
          "type": "number",
          "format": "int64"
        },
        "poolValues": {
          "type": "array",
          "maxItems": 5,
          "items": {
            "type": "number",
            "format": "int32"
          }
        }
      },
      "additionalProperties": false
    },
    "poolObjectAddResponse": {
      "type": "object",
      "required": [
        "status"
      ],
      "properties": {
        "status": {
          "type": "string"
        }
      }
    },
    "poolObjectQueryResponse": {
      "type": "object",
      "required": [
        "calculatedQuantile",
        "totalCount"
      ],
      "properties": {
        "calculatedQuantile": {
          "type": "number",
          "format": "double"
        },
        "totalCount": {
          "type": "number",
          "format": "int64"
        }
      }
    },
    "poolQueryRequest": {
      "type": "object",
      "required": [
        "poolId",
        "percentile"
      ],
      "properties": {
        "percentile": {
          "type": "number",
          "format": "double"
        },
        "poolId": {
          "type": "number",
          "format": "int64"
        }
      }
    },
    "poolQueryResponse": {
      "type": "object",
      "required": [
        "calculatedQuantile",
        "totalCount"
      ],
      "properties": {
        "calculatedQuantile": {
          "type": "number",
          "format": "double"
        },
        "totalCount": {
          "type": "number",
          "format": "int64"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Percentile",
    "version": "0.1.0"
  },
  "paths": {
    "/api/pools/add": {
      "post": {
        "description": "insert/append values to pools",
        "tags": [
          "pool"
        ],
        "operationId": "insertPool",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/poolObject"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/poolObjectAddResponse"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/errorResponse"
            }
          }
        }
      }
    },
    "/api/pools/query": {
      "post": {
        "description": "insert/append values to pools",
        "tags": [
          "pool"
        ],
        "operationId": "querryPool",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/poolQueryRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/poolObjectAddResponse"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/errorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "errorResponse": {
      "type": "string"
    },
    "poolObject": {
      "type": "object",
      "required": [
        "poolId",
        "poolValues"
      ],
      "properties": {
        "poolId": {
          "type": "number",
          "format": "int64"
        },
        "poolValues": {
          "type": "array",
          "maxItems": 5,
          "items": {
            "type": "number",
            "format": "int32"
          }
        }
      },
      "additionalProperties": false
    },
    "poolObjectAddResponse": {
      "type": "object",
      "required": [
        "status"
      ],
      "properties": {
        "status": {
          "type": "string"
        }
      }
    },
    "poolObjectQueryResponse": {
      "type": "object",
      "required": [
        "calculatedQuantile",
        "totalCount"
      ],
      "properties": {
        "calculatedQuantile": {
          "type": "number",
          "format": "double"
        },
        "totalCount": {
          "type": "number",
          "format": "int64"
        }
      }
    },
    "poolQueryRequest": {
      "type": "object",
      "required": [
        "poolId",
        "percentile"
      ],
      "properties": {
        "percentile": {
          "type": "number",
          "format": "double"
        },
        "poolId": {
          "type": "number",
          "format": "int64"
        }
      }
    },
    "poolQueryResponse": {
      "type": "object",
      "required": [
        "calculatedQuantile",
        "totalCount"
      ],
      "properties": {
        "calculatedQuantile": {
          "type": "number",
          "format": "double"
        },
        "totalCount": {
          "type": "number",
          "format": "int64"
        }
      }
    }
  }
}`))
}
