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
        "/quotes": {
            "get": {
                "description": "get quotes from all relayers.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quotes"
                ],
                "summary": "get quotes from all relayers.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "origin chain id to filter quotes by",
                        "name": "originChainID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "origin chain id to filter quotes by",
                        "name": "originTokenAddr",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "destination chain id to filter quotes by",
                        "name": "destChainID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "destination token address to filter quotes by",
                        "name": "destTokenAddr",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "relayer address to filter quotes by",
                        "name": "relayerAddr",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.GetQuoteResponse"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "upsert a quote from relayer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quotes"
                ],
                "summary": "get quotes from all relayers.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PutQuoteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.GetQuoteResponse": {
            "type": "object",
            "properties": {
                "dest_amount": {
                    "description": "DestAmount is the max amount of liquidity which exists for a given destination token, provided in the destination token decimals",
                    "type": "string"
                },
                "dest_chain_id": {
                    "description": "DestChainID is the chain which the relayer is willing to relay to",
                    "type": "integer"
                },
                "dest_fast_bridge_address": {
                    "description": "DestFastBridgeAddress is the address of the fast bridge contract on the destination chain",
                    "type": "string"
                },
                "dest_token_addr": {
                    "description": "DestToken is the token address for which the relayer willing to relay to",
                    "type": "string"
                },
                "fixed_fee": {
                    "description": "FixedFee is the fixed fee for the quote, provided in the destination token terms",
                    "type": "string"
                },
                "max_origin_amount": {
                    "description": "MaxOriginAmount is the maximum amount of origin tokens bridgeable",
                    "type": "string"
                },
                "origin_chain_id": {
                    "description": "OriginChainID is the chain which the relayer is willing to relay from",
                    "type": "integer"
                },
                "origin_fast_bridge_address": {
                    "description": "OriginFastBridgeAddress is the address of the fast bridge contract on the origin chain",
                    "type": "string"
                },
                "origin_token_addr": {
                    "description": "OriginTokenAddr is the token address for which the relayer willing to relay from",
                    "type": "string"
                },
                "relayer_addr": {
                    "description": "Address of the relayer providing the quote",
                    "type": "string"
                },
                "updated_at": {
                    "description": "UpdatedAt is the time that the quote was last upserted",
                    "type": "string"
                }
            }
        },
        "model.PutQuoteRequest": {
            "type": "object",
            "properties": {
                "dest_amount": {
                    "type": "string"
                },
                "dest_chain_id": {
                    "type": "integer"
                },
                "dest_fast_bridge_address": {
                    "type": "string"
                },
                "dest_token_addr": {
                    "type": "string"
                },
                "fixed_fee": {
                    "type": "string"
                },
                "max_origin_amount": {
                    "type": "string"
                },
                "origin_chain_id": {
                    "type": "integer"
                },
                "origin_fast_bridge_address": {
                    "type": "string"
                },
                "origin_token_addr": {
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