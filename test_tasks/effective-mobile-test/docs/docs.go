// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

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
        "/health": {
            "get": {
                "description": "Checking health of backend",
                "produces": [
                    "application/json"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Endpoint for getting all users",
                "produces": [
                    "application/json"
                ],
                "summary": "All users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/users.UserResponseDto"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Endpoint for creating and saving user to database",
                "produces": [
                    "application/json"
                ],
                "summary": "Create user",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/users.UserResponseDto"
                        }
                    }
                }
            }
        },
        "/users/health": {
            "get": {
                "description": "Checking health of users endpoint",
                "produces": [
                    "application/json"
                ],
                "summary": "Users Endpoint Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Endpoint for getting user with exact id",
                "produces": [
                    "application/json"
                ],
                "summary": "Get exact user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.UserResponseDto"
                        }
                    }
                }
            },
            "delete": {
                "description": "Endpoint for deleting user with exact id",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete exact user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/users.UserResponseDto"
                        }
                    }
                }
            },
            "patch": {
                "description": "Endpoint for updating user with exact id",
                "produces": [
                    "application/json"
                ],
                "summary": "Update exact user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/users.UserResponseDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "users.UserResponseDto": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "second_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Documentation",
	Description:      "Effective Mobile test task in Gin Framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}