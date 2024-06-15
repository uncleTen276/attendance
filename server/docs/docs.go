// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Nguyen Tri",
            "email": "tringuyen2762001@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/attendance/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Attendance"
                ],
                "summary": "api use to get user attendance record",
                "responses": {
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "check-in time should be format : 2024-06-15T08:00:00Z",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Attendance"
                ],
                "summary": "api use to update attendaceRecord",
                "parameters": [
                    {
                        "description": "update record",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AttendaceUpdate"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "check-in time should be format : 2024-06-15T08:00:00Z",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Attendance"
                ],
                "summary": "api use for employee checkIn",
                "parameters": [
                    {
                        "description": "Check In position",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AttendanceCheckIn"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "use for login response the access_Token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserToken"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "create new user",
                "parameters": [
                    {
                        "description": "New User",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/me": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "route get user Id from token then get user information",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get user information by Id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AttendaceUpdate": {
            "type": "object",
            "required": [
                "check_in",
                "position",
                "record_id"
            ],
            "properties": {
                "check_in": {
                    "type": "string"
                },
                "detail": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "record_id": {
                    "type": "integer"
                }
            }
        },
        "models.AttendanceCheckIn": {
            "type": "object",
            "required": [
                "check_in",
                "position"
            ],
            "properties": {
                "check_in": {
                    "type": "string"
                },
                "detail": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                }
            }
        },
        "models.CreateUser": {
            "type": "object",
            "required": [
                "firstName",
                "identifier",
                "lastName",
                "password"
            ],
            "properties": {
                "firstName": {
                    "type": "string",
                    "maxLength": 50
                },
                "identifier": {
                    "type": "string",
                    "maxLength": 100
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 50
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "identifier": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                }
            }
        },
        "models.UserLogin": {
            "type": "object",
            "required": [
                "identifier",
                "password"
            ],
            "properties": {
                "identifier": {
                    "type": "string",
                    "maxLength": 100
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8
                }
            }
        },
        "models.UserToken": {
            "type": "object",
            "properties": {
                "access_Token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
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
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Fiber Go API",
	Description:      "employee attendance",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
