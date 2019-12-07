// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-07 23:13:46.910114 +0800 CST m=+0.040049856

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/ping": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试"
                ],
                "summary": "接口名称",
                "responses": {
                    "200": {},
                    "500": {}
                }
            }
        },
        "/api/ping1": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试时"
                ],
                "summary": "接口名称",
                "responses": {
                    "200": {},
                    "500": {}
                }
            }
        },
        "/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "注册",
                "responses": {
                    "200": {},
                    "500": {}
                }
            }
        },
        "/user": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "接口名称",
                "responses": {
                    "200": {},
                    "500": {}
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
