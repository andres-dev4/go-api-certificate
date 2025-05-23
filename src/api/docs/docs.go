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
        "/": {
            "get": {
                "description": "Devuelve un Hola Mundo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ejemplo"
                ],
                "summary": "Endpoint de ejemplo",
                "responses": {
                    "200": {
                        "description": "Hola Mundo!",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/certificates": {
            "get": {
                "description": "Get all certificates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificates"
                ],
                "summary": "List all certificates",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Certificate"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Certificate": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "credential": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "expiry_date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "issue_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Certificate API",
	Description:      "API para gestión de certificados",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
