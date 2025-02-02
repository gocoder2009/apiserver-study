// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "lkong",
            "url": "http://www.swagger.io/support",
            "email": "466701708@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Login generates the authentication token",
                "parameters": [
                    {
                        "description": "Password",
                        "name": "LoginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ\"}}",
                        "schema": {
                            "$ref": "#/definitions/user.LoginResponse"
                        }
                    }
                }
            }
        },
        "/sd/cpu": {
            "get": {
                "description": "Checks the cpu usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Checks the cpu usage",
                "responses": {
                    "200": {
                        "description": "CRITICAL - Load average: 1.78, 1.99, 2.02 | Cores: 2",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sd/disk": {
            "get": {
                "description": "Checks the disk usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Checks the disk usage",
                "responses": {
                    "200": {
                        "description": "OK - Free space: 17233MB (16GB) / 51200MB (50GB) | Used: 33%",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sd/health": {
            "get": {
                "description": "Shows OK as the ping-pong result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Shows OK as the ping-pong result",
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
        "/sd/ram": {
            "get": {
                "description": "Checks the ram usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Checks the ram usage",
                "responses": {
                    "200": {
                        "description": "OK - Free space: 402MB (0GB) / 8192MB (8GB) | Used: 4%",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "get users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "List the users in the database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"totalCount\":1,\"userList\":[{\"id\":0,\"username\":\"admin\",\"random\":\"user\t'admin'\tget\trandom\tstring\t'EnqntiSig'\",\"password\":\"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG\",\"createdAt\":\"2018-05-28\t00:25:33\",\"updatedAt\":\"2018-05-28\t00:25:33\"}]}}",
                        "schema": {
                            "$ref": "#/definitions/user.SwaggerListResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Add a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Add new user to the database",
                "parameters": [
                    {
                        "description": "Create a new user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"username\":\"kong\"}}",
                        "schema": {
                            "$ref": "#/definitions/user.CreateResponse"
                        }
                    }
                }
            }
        },
        "/v1/user/{id}": {
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Update a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a user info by the user identifier",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The user's database id index num",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":null}",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete a user by the user identifier",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The user's database id index num",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":null}",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/v1/user/{username}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get an user by username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get an user by the user identifier",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"username\":\"kong\",\"password\":\"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS\"}}",
                        "schema": {
                            "$ref": "#/definitions/model.UserModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "sayHello": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.UserModel": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 128,
                    "minLength": 5
                },
                "username": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 1
                }
            }
        },
        "user.CreateRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.CreateResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "user.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "user.SwaggerListResponse": {
            "type": "object",
            "properties": {
                "totalCount": {
                    "type": "integer"
                },
                "userList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UserInfo"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "description": "请在token字符串前加` + "`" + `Bearer+空格` + "`" + `",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Apiserver Example API",
	Description:      "apiserver demo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
