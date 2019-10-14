package api

func init() {
	Swagger.Add("data_feed", `{
  "swagger": "2.0",
  "info": {
    "title": "api/external/data_feed/data_feed.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/datafeed/destination": {
      "post": {
        "operationId": "AddDestination",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/datafeedAddDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/datafeedAddDestinationRequest"
            }
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      }
    },
    "/datafeed/destination/{id}": {
      "get": {
        "operationId": "GetDestination",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/datafeedGetDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      },
      "delete": {
        "operationId": "DeleteDestination",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/datafeedDeleteDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      },
      "patch": {
        "operationId": "UpdateDestination",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/datafeedUpdateDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/datafeedUpdateDestinationRequest"
            }
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      }
    },
    "/datafeed/destinations": {
      "post": {
        "operationId": "ListDestinations",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/datafeedListDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/datafeedListDestinationRequest"
            }
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      }
    },
    "/datafeed/destinations/test": {
      "post": {
        "operationId": "TestDestination",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/datafeedTestDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/datafeedURLValidationRequest"
            }
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      }
    }
  },
  "definitions": {
    "datafeedAddDestinationRequest": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        }
      }
    },
    "datafeedAddDestinationResponse": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "datafeedDeleteDestinationResponse": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "datafeedGetDestinationResponse": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean",
          "format": "boolean"
        },
        "id": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        }
      }
    },
    "datafeedListDestinationRequest": {
      "type": "object"
    },
    "datafeedListDestinationResponse": {
      "type": "object",
      "properties": {
        "destinations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/datafeedGetDestinationResponse"
          }
        }
      }
    },
    "datafeedSecretId": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "datafeedTestDestinationResponse": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "datafeedURLValidationRequest": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string"
        },
        "username_password": {
          "$ref": "#/definitions/datafeedUsernamePassword"
        },
        "secret_id": {
          "$ref": "#/definitions/datafeedSecretId"
        }
      }
    },
    "datafeedUpdateDestinationRequest": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        }
      }
    },
    "datafeedUpdateDestinationResponse": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "datafeedUsernamePassword": {
      "type": "object",
      "properties": {
        "username": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    }
  }
}
`)
}
