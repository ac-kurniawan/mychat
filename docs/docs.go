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
        "/room_chat": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roomChat"
                ],
                "summary": "get room chat by session ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "session id",
                        "name": "session_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "filter session status",
                        "name": "session_status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.RoomChatDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.HttpErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.HttpErrorModel"
                        }
                    }
                }
            }
        },
        "/room_chats": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roomChat"
                ],
                "summary": "get room chat list by participant ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "participant id",
                        "name": "participant",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "filter session status",
                        "name": "session_status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.RoomChatDto"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.HttpErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.HttpErrorModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ChatDto": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "messageType": {
                    "type": "string"
                },
                "readAt": {
                    "type": "string"
                },
                "replyforChatId": {
                    "type": "string"
                },
                "senderId": {
                    "type": "string"
                },
                "sessionId": {
                    "type": "string"
                }
            }
        },
        "dto.RoomChatDto": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "participantGroup": {
                    "type": "string"
                },
                "sessionChats": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.SessionChatDto"
                    }
                }
            }
        },
        "dto.SessionChatDto": {
            "type": "object",
            "properties": {
                "chats": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ChatDto"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "endedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "roomChatId": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "http.HttpErrorModel": {
            "type": "object",
            "properties": {
                "code": {},
                "httpStatus": {
                    "type": "integer"
                },
                "message": {
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
