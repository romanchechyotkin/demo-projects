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
        "/auth/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_controller_v1_request.Login"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_controller_v1_response.Login"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Registration",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_controller_v1_request.Registration"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_controller_v1_response.Registration"
                        }
                    }
                }
            }
        },
        "/v1/flat/create": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create Flat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flat"
                ],
                "summary": "Create Flat",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_controller_v1_request.CreateFlat"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_controller_v1_response.Flat"
                        }
                    }
                }
            }
        },
        "/v1/flat/update": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Update Flat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flat"
                ],
                "summary": "Update Flat",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_controller_v1_request.UpdateFlat"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_controller_v1_response.Flat"
                        }
                    }
                }
            }
        },
        "/v1/house/create": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create House",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "house"
                ],
                "summary": "Create House",
                "parameters": [
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_controller_v1_request.CreateHouse"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_controller_v1_response.House"
                        }
                    }
                }
            }
        },
        "/v1/house/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get House Flats",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "house"
                ],
                "summary": "Get House Flats",
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
                            "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_controller_v1_response.HouseFlats"
                        }
                    }
                }
            }
        },
        "/v1/house/{id}/subscribe": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Subscribe for house updates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "house"
                ],
                "summary": "Subscribe for house updates",
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
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_romanchechyotkin_avito_test_task_internal_controller_v1_request.CreateFlat": {
            "type": "object"
        },
        "github_com_romanchechyotkin_avito_test_task_internal_controller_v1_request.CreateHouse": {
            "type": "object",
            "required": [
                "address",
                "year"
            ],
            "properties": {
                "address": {
                    "type": "string",
                    "example": "ул. Новая, д. 1"
                },
                "developer": {
                    "type": "string",
                    "example": "ООО Компания"
                },
                "year": {
                    "type": "integer",
                    "example": 2022
                }
            }
        },
        "github_com_romanchechyotkin_avito_test_task_internal_controller_v1_request.Login": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 4,
                    "example": "password"
                }
            }
        },
        "github_com_romanchechyotkin_avito_test_task_internal_controller_v1_request.Registration": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 4,
                    "example": "password"
                },
                "user_type": {
                    "type": "string",
                    "enum": [
                        "client",
                        "moderator"
                    ],
                    "example": "client"
                }
            }
        },
        "github_com_romanchechyotkin_avito_test_task_internal_controller_v1_request.UpdateFlat": {
            "type": "object",
            "required": [
                "id",
                "status"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 123
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "created",
                        "approved",
                        "declined",
                        "on moderation"
                    ],
                    "example": "on moderation"
                }
            }
        },
        "github_com_romanchechyotkin_avito_test_task_internal_controller_v1_response.Flat": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-08-09T00:00:00Z"
                },
                "house_id": {
                    "type": "integer",
                    "example": 1
                },
                "id": {
                    "type": "integer",
                    "example": 123
                },
                "moderation_status": {
                    "type": "string",
                    "example": "created"
                },
                "number": {
                    "type": "integer",
                    "example": 1
                },
                "price": {
                    "type": "integer",
                    "example": 123
                },
                "rooms_amount": {
                    "type": "integer",
                    "example": 4
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-08-09T00:00:00Z"
                }
            }
        },
        "github_com_romanchechyotkin_avito_test_task_internal_controller_v1_response.House": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "Улица Пушкина 1"
                },
                "created_at": {
                    "type": "string",
                    "example": "2024-08-09T00:00:00Z"
                },
                "developer": {
                    "type": "string",
                    "example": "ООО Компания"
                },
                "id": {
                    "type": "integer",
                    "example": 123
                },
                "updated_at": {
                    "type": "string",
                    "example": "2024-08-09T00:00:00Z"
                },
                "year": {
                    "type": "integer",
                    "example": 1999
                }
            }
        },
        "github_com_romanchechyotkin_avito_test_task_internal_controller_v1_response.HouseFlats": {
            "type": "object",
            "properties": {
                "flats": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_romanchechyotkin_avito_test_task_internal_entity.Flat"
                    }
                }
            }
        },
        "github_com_romanchechyotkin_avito_test_task_internal_controller_v1_response.Login": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "auth token"
                }
            }
        },
        "github_com_romanchechyotkin_avito_test_task_internal_controller_v1_response.Registration": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string",
                    "example": "cae36e0f-69e5-4fa8-a179-a52d083c5549"
                }
            }
        },
        "github_com_romanchechyotkin_avito_test_task_internal_entity.Flat": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "houseID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "moderationStatus": {
                    "type": "string"
                },
                "number": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "roomsAmount": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
