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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a Plantbook project API description. You can find out more about us at [https://github.com/kaatinga/plantbook](https://github.com/kaatinga/plantbook).",
    "title": "Plantbook",
    "license": {
      "name": "MIT",
      "url": "https://opensource.org/licenses/mit-license.php"
    },
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/api/v1/refplant/{id}": {
      "get": {
        "description": "Returns a single reference plant",
        "produces": [
          "application/json"
        ],
        "tags": [
          "refplant"
        ],
        "summary": "Find reference plant by ID",
        "operationId": "getRefPlantById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of reference plant to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/RefPlant"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Plant not found"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "error"
              }
            }
          }
        }
      }
    },
    "/api/v1/refplants": {
      "get": {
        "description": "find reference plants by parameters or all",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "refplant"
        ],
        "summary": "find reference plants",
        "operationId": "getRefPlants",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "plant category",
            "name": "category",
            "in": "query"
          },
          {
            "type": "string",
            "name": "kind",
            "in": "query"
          },
          {
            "type": "string",
            "name": "recommendPosition",
            "in": "query"
          },
          {
            "type": "string",
            "name": "regardToLight",
            "in": "query"
          },
          {
            "type": "string",
            "name": "regardToMoisture",
            "in": "query"
          },
          {
            "type": "string",
            "name": "floweringTime",
            "in": "query"
          },
          {
            "type": "string",
            "name": "hight",
            "in": "query"
          },
          {
            "type": "string",
            "name": "classifiers",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "limit",
            "in": "query",
            "required": true,
            "allowEmptyValue": true
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "offset",
            "in": "query",
            "required": true,
            "allowEmptyValue": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/RefPlant"
              }
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "404": {
            "description": "Plants not found"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "error"
              }
            }
          }
        }
      }
    },
    "/api/v1/user": {
      "post": {
        "description": "This can only be done by the logged in user.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Create user object without id and status",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created user with id and status",
            "schema": {
              "$ref": "#/definitions/User"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/api/v1/user/login": {
      "post": {
        "description": "Logins user by passing login/password",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logins user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "name": "login/password",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/UserLoginPassword"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "headers": {
              "Set-Cookie": {
                "type": "string",
                "description": "set cookie with jwt token value with name plantbook_token"
              },
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/api/v1/user/logout": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "200": {
            "description": "successful operation, set expired cookie",
            "headers": {
              "Set-Cookie": {
                "type": "string",
                "description": "set expired cookie with name plantbook_token"
              },
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/api/v1/user/{username}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get user by user name",
        "operationId": "getUserByName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched. Use user1 for testing. ",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "put": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Updated user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "name that need to be updated",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid user supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be deleted",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      }
    },
    "/api/v1/version": {
      "get": {
        "description": "Shows api version.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "health"
        ],
        "summary": "Get app version",
        "operationId": "apiVersion",
        "responses": {
          "200": {
            "description": "version info",
            "schema": {
              "$ref": "#/definitions/ApiVersion"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/health/live": {
      "get": {
        "description": "Checks web service is working",
        "tags": [
          "health"
        ],
        "summary": "Probe service alive",
        "operationId": "healthAlive",
        "responses": {
          "200": {
            "description": "successful operation, service alive",
            "schema": {
              "type": "string"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/health/ready": {
      "get": {
        "description": "Checks web service is working, and subservices available",
        "tags": [
          "health"
        ],
        "summary": "Probe service ready",
        "operationId": "healthReady",
        "responses": {
          "200": {
            "description": "successful operation, service ready to process requests",
            "schema": {
              "type": "string"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/metrics": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "tags": [
          "health"
        ],
        "summary": "Prometheus metrics",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "ApiVersion": {
      "type": "object",
      "properties": {
        "build_at": {
          "type": "string",
          "format": "date-time"
        },
        "githash": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    },
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Info": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "RefPlant": {
      "type": "object",
      "required": [
        "title",
        "images"
      ],
      "properties": {
        "category": {
          "type": "integer",
          "format": "int32"
        },
        "createdAt": {
          "x-go-type": {
            "hints": {
              "noValidation": true
            },
            "import": {
              "package": "time"
            },
            "type": "Time"
          }
        },
        "creater": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "images": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "infos": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Info"
          }
        },
        "modifiedAt": {
          "x-go-type": {
            "hints": {
              "noValidation": true
            },
            "import": {
              "package": "time"
            },
            "type": "Time"
          },
          "x-nullable": true
        },
        "modifier": {
          "type": "string"
        },
        "shortInfo": {
          "$ref": "#/definitions/ShortInfo"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "ShortInfo": {
      "type": "object",
      "properties": {
        "classifiers": {
          "type": "string"
        },
        "flowering_Time": {
          "type": "string"
        },
        "hight": {
          "type": "string"
        },
        "kind": {
          "type": "string"
        },
        "recommend_Position": {
          "type": "string"
        },
        "regard_To_Light": {
          "type": "string"
        },
        "regard_To_Moisture": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "lastName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "phone": {
          "type": "string"
        },
        "userRole": {
          "type": "integer",
          "format": "int32",
          "maximum": 2,
          "minimum": 1
        },
        "userStatus": {
          "description": "User Status",
          "type": "integer",
          "format": "int32"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "UserLoginPassword": {
      "type": "object",
      "required": [
        "login",
        "password"
      ],
      "properties": {
        "login": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything about your Plants",
      "name": "plant"
    },
    {
      "description": "Operations about user",
      "name": "user"
    },
    {
      "description": "Operations about health checks",
      "name": "health"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a Plantbook project API description. You can find out more about us at [https://github.com/kaatinga/plantbook](https://github.com/kaatinga/plantbook).",
    "title": "Plantbook",
    "license": {
      "name": "MIT",
      "url": "https://opensource.org/licenses/mit-license.php"
    },
    "version": "1.0.0"
  },
  "basePath": "/",
  "paths": {
    "/api/v1/refplant/{id}": {
      "get": {
        "description": "Returns a single reference plant",
        "produces": [
          "application/json"
        ],
        "tags": [
          "refplant"
        ],
        "summary": "Find reference plant by ID",
        "operationId": "getRefPlantById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of reference plant to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/RefPlant"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Plant not found"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "error"
              }
            }
          }
        }
      }
    },
    "/api/v1/refplants": {
      "get": {
        "description": "find reference plants by parameters or all",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "refplant"
        ],
        "summary": "find reference plants",
        "operationId": "getRefPlants",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "description": "plant category",
            "name": "category",
            "in": "query"
          },
          {
            "type": "string",
            "name": "kind",
            "in": "query"
          },
          {
            "type": "string",
            "name": "recommendPosition",
            "in": "query"
          },
          {
            "type": "string",
            "name": "regardToLight",
            "in": "query"
          },
          {
            "type": "string",
            "name": "regardToMoisture",
            "in": "query"
          },
          {
            "type": "string",
            "name": "floweringTime",
            "in": "query"
          },
          {
            "type": "string",
            "name": "hight",
            "in": "query"
          },
          {
            "type": "string",
            "name": "classifiers",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "limit",
            "in": "query",
            "required": true,
            "allowEmptyValue": true
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "offset",
            "in": "query",
            "required": true,
            "allowEmptyValue": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/RefPlant"
              }
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "404": {
            "description": "Plants not found"
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "error"
              }
            }
          }
        }
      }
    },
    "/api/v1/user": {
      "post": {
        "description": "This can only be done by the logged in user.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Create user object without id and status",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "created user with id and status",
            "schema": {
              "$ref": "#/definitions/User"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/api/v1/user/login": {
      "post": {
        "description": "Logins user by passing login/password",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logins user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "name": "login/password",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/UserLoginPassword"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "headers": {
              "Set-Cookie": {
                "type": "string",
                "description": "set cookie with jwt token value with name plantbook_token"
              },
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/api/v1/user/logout": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "200": {
            "description": "successful operation, set expired cookie",
            "headers": {
              "Set-Cookie": {
                "type": "string",
                "description": "set expired cookie with name plantbook_token"
              },
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/api/v1/user/{username}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get user by user name",
        "operationId": "getUserByName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched. Use user1 for testing. ",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "put": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Updated user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "name that need to be updated",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid user supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be deleted",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      }
    },
    "/api/v1/version": {
      "get": {
        "description": "Shows api version.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "health"
        ],
        "summary": "Get app version",
        "operationId": "apiVersion",
        "responses": {
          "200": {
            "description": "version info",
            "schema": {
              "$ref": "#/definitions/ApiVersion"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/health/live": {
      "get": {
        "description": "Checks web service is working",
        "tags": [
          "health"
        ],
        "summary": "Probe service alive",
        "operationId": "healthAlive",
        "responses": {
          "200": {
            "description": "successful operation, service alive",
            "schema": {
              "type": "string"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/health/ready": {
      "get": {
        "description": "Checks web service is working, and subservices available",
        "tags": [
          "health"
        ],
        "summary": "Probe service ready",
        "operationId": "healthReady",
        "responses": {
          "200": {
            "description": "successful operation, service ready to process requests",
            "schema": {
              "type": "string"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          }
        }
      }
    },
    "/metrics": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "tags": [
          "health"
        ],
        "summary": "Prometheus metrics",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "ApiVersion": {
      "type": "object",
      "properties": {
        "build_at": {
          "type": "string",
          "format": "date-time"
        },
        "githash": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    },
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Info": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "RefPlant": {
      "type": "object",
      "required": [
        "title",
        "images"
      ],
      "properties": {
        "category": {
          "type": "integer",
          "format": "int32"
        },
        "createdAt": {
          "x-go-type": {
            "hints": {
              "noValidation": true
            },
            "import": {
              "package": "time"
            },
            "type": "Time"
          }
        },
        "creater": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "images": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "infos": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Info"
          }
        },
        "modifiedAt": {
          "x-go-type": {
            "hints": {
              "noValidation": true
            },
            "import": {
              "package": "time"
            },
            "type": "Time"
          },
          "x-nullable": true
        },
        "modifier": {
          "type": "string",
          "minLength": 0
        },
        "shortInfo": {
          "$ref": "#/definitions/ShortInfo"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "ShortInfo": {
      "type": "object",
      "properties": {
        "classifiers": {
          "type": "string"
        },
        "flowering_Time": {
          "type": "string"
        },
        "hight": {
          "type": "string"
        },
        "kind": {
          "type": "string"
        },
        "recommend_Position": {
          "type": "string"
        },
        "regard_To_Light": {
          "type": "string"
        },
        "regard_To_Moisture": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "lastName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "phone": {
          "type": "string"
        },
        "userRole": {
          "type": "integer",
          "format": "int32",
          "maximum": 2,
          "minimum": 1
        },
        "userStatus": {
          "description": "User Status",
          "type": "integer",
          "format": "int32"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "UserLoginPassword": {
      "type": "object",
      "required": [
        "login",
        "password"
      ],
      "properties": {
        "login": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything about your Plants",
      "name": "plant"
    },
    {
      "description": "Operations about user",
      "name": "user"
    },
    {
      "description": "Operations about health checks",
      "name": "health"
    }
  ]
}`))
}
