// Package swagger GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

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
        "termsOfService": "https://github.com/swaggo/swag",
        "contact": {
            "name": "zhangsan"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/answer/create": {
            "post": {
                "description": "创建回答",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qa"
                ],
                "summary": "创建回答",
                "parameters": [
                    {
                        "description": "创建回答参数",
                        "name": "answerCreateParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/qa.answerCreateParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/answer/delete": {
            "get": {
                "description": "创建回答",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qa"
                ],
                "summary": "创建回答",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "删除id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/question/create": {
            "post": {
                "description": "创建问题",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qa"
                ],
                "summary": "创建问题",
                "parameters": [
                    {
                        "description": "创建问题参数",
                        "name": "questionCreateParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/qa.questionCreateParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/question/delete": {
            "get": {
                "description": "删除问题，同时删除问题中的所有答案",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qa"
                ],
                "summary": "删除问题",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "删除id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/question/detail": {
            "get": {
                "description": "获取问题详情，包括问题的所有回答",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qa"
                ],
                "summary": "获取问题详细",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "问题id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "问题详情，带回答和作者",
                        "schema": {
                            "type": "QuestionDTO"
                        }
                    }
                }
            }
        },
        "/api/question/edit": {
            "post": {
                "description": "编辑问题",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qa"
                ],
                "summary": "编辑问题",
                "parameters": [
                    {
                        "description": "编辑问题参数",
                        "name": "questionEditParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/qa.questionEditParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/question/list": {
            "get": {
                "description": "获取问题列表, 包含作者信息, 不包含回答",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "qa"
                ],
                "summary": "获取问题列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "列表页页数",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "列表页单页个数",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "问题列表",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/qa.QuestionDTO"
                            }
                        }
                    }
                }
            }
        },
        "/api/user/login": {
            "post": {
                "description": "用户登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "login with param",
                        "name": "loginParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.loginParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/logout": {
            "get": {
                "description": "调用表示用户登出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户登出",
                "responses": {
                    "200": {
                        "description": "用户登出成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/register": {
            "post": {
                "description": "用户注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "注册参数",
                        "name": "registerParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.registerParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注册成功",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/register/verify": {
            "get": {
                "description": "使用token验证用户注册信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "验证注册信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "注册token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注册成功，请进入登录页面",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "qa.AnswerDTO": {
            "type": "object",
            "properties": {
                "author": {
                    "description": "作者",
                    "$ref": "#/definitions/user.UserDTO"
                },
                "context": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "qa.QuestionDTO": {
            "type": "object",
            "properties": {
                "answer_num": {
                    "type": "integer"
                },
                "answers": {
                    "description": "回答",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/qa.AnswerDTO"
                    }
                },
                "author": {
                    "description": "作者",
                    "$ref": "#/definitions/user.UserDTO"
                },
                "context": {
                    "description": "在列表页，只显示前200个字符",
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "qa.answerCreateParam": {
            "type": "object",
            "required": [
                "Context",
                "question_id"
            ],
            "properties": {
                "Context": {
                    "type": "string"
                },
                "question_id": {
                    "type": "integer"
                }
            }
        },
        "qa.questionCreateParam": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "qa.questionEditParam": {
            "type": "object",
            "required": [
                "content",
                "id",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "user.UserDTO": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "user.loginParam": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.registerParam": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
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
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "bbs 论坛",
	Description: "hade测试",
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
	swag.Register(swag.Name, &s{})
}
