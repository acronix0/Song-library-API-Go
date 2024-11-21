// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://example.com/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://example.com/support",
            "email": "support@example.com"
        },
        "license": {
            "name": "MIT",
            "url": "http://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/songs": {
            "get": {
                "description": "Get a paginated list of songs.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Get Songs",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Number of records to skip",
                        "name": "skip",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Number of records to fetch",
                        "name": "take",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_acronix0_song-libary-api_internal_dto.SongDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid query parameters",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new song to the library.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Add Song",
                "parameters": [
                    {
                        "description": "Song object",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_acronix0_song-libary-api_internal_dto.SongDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_acronix0_song-libary-api_internal_dto.SongDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a song by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Delete Song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the song to delete",
                        "name": "song_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid song_id parameter",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    }
                }
            }
        },
        "/songs/text": {
            "get": {
                "description": "Get the text of a song by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Get Song Text",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the song",
                        "name": "song_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Number of verses to skip",
                        "name": "skip",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Number of verses to fetch",
                        "name": "take",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "string",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid query parameters",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    },
                    "404": {
                        "description": "Song text not found",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    }
                }
            }
        },
        "/songs/{id}": {
            "put": {
                "description": "Update an existing song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Update song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated song data",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_acronix0_song-libary-api_internal_dto.SongDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_acronix0_song-libary-api_internal_dto.SongDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_delivery_http_v1.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_acronix0_song-libary-api_internal_dto.SongDTO": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "lyrics": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "song_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "internal_delivery_http_v1.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Song Library API",
	Description:      "API for managing songs and their texts.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
