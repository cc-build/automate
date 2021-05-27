package api

func init() {
	Swagger.Add("user_settings", `{
  "swagger": "2.0",
  "info": {
    "title": "external/user_settings/user_settings.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v0/user-settings/{user.name}/{user.connector}": {
      "get": {
        "summary": "GetUserSettings returns all of the preferences for a given user",
        "operationId": "UserSettingsService_GetUserSettings",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.user_settings.GetUserSettingsResponse"
            }
          },
          "default": {
            "description": "An unexpected error response",
            "schema": {
              "$ref": "#/definitions/grpc.gateway.runtime.Error"
            }
          }
        },
        "parameters": [
          {
            "name": "user.name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "user.connector",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "UserSettingsService"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.user_settings.GetUserSettingsResponse": {
      "type": "object",
      "properties": {
        "user": {
          "$ref": "#/definitions/chef.automate.api.user_settings.User"
        },
        "settings": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/chef.automate.api.user_settings.UserSettingValue"
          }
        }
      }
    },
    "chef.automate.api.user_settings.User": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "connector": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.user_settings.UserSettingValue": {
      "type": "object",
      "properties": {
        "default_value": {
          "type": "string",
          "description": "Default value for this setting."
        },
        "value": {
          "type": "string",
          "description": "Value for this setting."
        },
        "enabled": {
          "type": "boolean",
          "format": "boolean",
          "title": "Enabled"
        }
      }
    },
    "google.protobuf.Any": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpc.gateway.runtime.Error": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/google.protobuf.Any"
          }
        }
      }
    }
  }
}
`)
}
