// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/healthcheck": {
            "get": {
                "security": [
                    {
                        "AuthenticationToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show the current status of the server by running the stored command",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.HealthCheckResponse"
                        }
                    },
                    "401": {
                        "description": "Request did not provide a valid authentication token"
                    },
                    "403": {
                        "description": "Request did not come from an IP within the whitelisted IP ranges"
                    }
                }
            }
        },
        "/status": {
            "get": {
                "security": [
                    {
                        "AuthenticationToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show the current owner of the server that the agent is currently running on",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.StatusCheckResponse"
                        }
                    },
                    "401": {
                        "description": "Request did not provide a valid authentication token"
                    },
                    "403": {
                        "description": "Request did not come from an IP within the whitelisted IP ranges"
                    }
                }
            }
        }
    },
    "definitions": {
        "main.HealthCheckData": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "integer"
                },
                "stderr": {
                    "type": "string"
                },
                "stdout": {
                    "type": "string"
                }
            }
        },
        "main.HealthCheckResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/main.HealthCheckData"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "main.StatusCheckData": {
            "type": "object",
            "properties": {
                "identifier": {
                    "type": "string"
                }
            }
        },
        "main.StatusCheckResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/main.StatusCheckData"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "AuthenticationToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
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
	Title:       "CTFd King of the Hill Agent",
	Description: "This agent implements a small HTTP interface for scoring servers (i.e. CTFd Enterprise) to poll during a King of the Hill CTF.",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
