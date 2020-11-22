// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/payment/deposit": {
            "post": {
                "description": "Transferring money between wallets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment System"
                ],
                "summary": "transfer money",
                "parameters": [
                    {
                        "description": "Request Payload",
                        "name": "Payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TransferRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success operation",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/service.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.TransferResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/payment/wallet": {
            "get": {
                "description": "get wallet by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment System"
                ],
                "summary": "Get wallet",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Wallet ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success operation",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/service.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.WalletResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "create new wallet, return last created id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment System"
                ],
                "summary": "New wallet",
                "parameters": [
                    {
                        "description": "Request Payload",
                        "name": "Payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewWalletRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success operation",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/service.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.WalletResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.DepositRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "wallet_id": {
                    "type": "integer"
                }
            }
        },
        "dto.DepositResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "wallet_id": {
                    "type": "integer"
                }
            }
        },
        "dto.NewWalletRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.TransferRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "wallet_from": {
                    "type": "integer"
                },
                "wallet_to": {
                    "type": "integer"
                }
            }
        },
        "dto.TransferResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "wallet_from": {
                    "type": "integer"
                }
            }
        },
        "dto.WalletResponse": {
            "type": "object",
            "properties": {
                "Name": {
                    "type": "string"
                },
                "balance": {
                    "type": "number"
                },
                "created_at": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "integer"
                }
            }
        },
        "service.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "error": {
                    "type": "string"
                },
                "error_detail": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "description": "Group error code",
                            "type": "string",
                            "example": "1.1.1"
                        },
                        "error_origin": {
                            "description": "Origin of error (group)",
                            "type": "string",
                            "example": "invalid parameter"
                        },
                        "extra": {
                            "description": "Extra fields",
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        },
                        "id": {
                            "description": "Error Id in current request",
                            "type": "string",
                            "example": "1dQqPlQgJuPPJJfAd7pjmfBWMoP"
                        }
                    }
                },
                "request_id": {
                    "description": "The X-Request-ID from request header. The request ID represented in the HTTP header X-Request-ID let you to link all the log lines which are common to a single web request.",
                    "type": "string",
                    "example": "948b9acf-36c0-452d-af21-66b362778fa3"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Payment System API",
	Description: "Backend API Payment System",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
