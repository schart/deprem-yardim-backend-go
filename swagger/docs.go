// Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/events": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Create Event areas with request body",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/feeds/areas": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feed"
                ],
                "summary": "Get Feed areas with query strings",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Sw Lat",
                        "name": "sw_lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Sw Lng",
                        "name": "sw_lng",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Ne Lat",
                        "name": "ne_lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Ne Lng",
                        "name": "ne_lng",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Timestamp",
                        "name": "time_stamp",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/feeds/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feed"
                ],
                "summary": "Get Feeds with given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Feed Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Healthcheck"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.request": {
            "type": "object",
            "properties": {
                "channel": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "payload": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:80",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "IT Afet Yardım",
	Description:      "Afet Harita API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
