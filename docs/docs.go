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
        "/geocode": {
            "get": {
                "description": "Get geocoding for a given city",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geocode"
                ],
                "summary": "Get geocoding",
                "parameters": [
                    {
                        "type": "string",
                        "description": "City",
                        "name": "city",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/geocode/reverse": {
            "get": {
                "description": "Get city from latitude and longitude",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geocode"
                ],
                "summary": "Get city",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Latitude",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Longitude",
                        "name": "long",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/weather/forecast": {
            "get": {
                "description": "Get 5-day forecast for a given latitude and longitude",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "weather"
                ],
                "summary": "Get 5-day forecast",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Latitude",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Longitude",
                        "name": "long",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/weather/now": {
            "get": {
                "description": "Get current weather for a given latitude and longitude",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "weather"
                ],
                "summary": "Get current weather",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Latitude",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Longitude",
                        "name": "long",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8181",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Weather Wrapper API",
	Description:      "This is a wrapper for the OpenWeatherMap API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
