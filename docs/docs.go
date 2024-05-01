// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/Yhelee99",
        "contact": {
            "name": "Yhelee",
            "url": "https://github.com/Yhelee99",
            "email": "yhelee99@gmail.com"
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
        "/community": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取社区列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区相关接口"
                ],
                "summary": "获取社区列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseCommunity"
                        }
                    }
                }
            }
        },
        "/community/:id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据社区id查询社区详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区相关接口"
                ],
                "summary": "获取社区列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "社区ID,可不传,不传获取全量",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseCommunity"
                        }
                    }
                }
            }
        },
        "/createpost": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建帖子接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "创建帖子",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mod.Post"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/getpostslist": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "可选排序方式的获取帖子列表的接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "获取帖子列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "获取帖子的参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mod.ParamsGetPostList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePost"
                        }
                    }
                }
            }
        },
        "/post/:id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据postid获取帖子详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "获取帖子详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "帖子id",
                        "name": "postid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePost"
                        }
                    }
                }
            }
        },
        "/post/voted": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "帖子投票接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "帖子投票",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "获取帖子参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mod.PostVoted"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取帖子列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "获取帖子列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页数",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePost"
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mod.ParamSignIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseUser"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "用户注册接口",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mod.ParamSignUp"
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
        "controller.ResCode": {
            "type": "integer",
            "enum": [
                10000,
                10001,
                10002,
                10003,
                10004,
                10005,
                10006,
                10007
            ],
            "x-enum-varnames": [
                "SuccessCode",
                "ErrorCodeInvalidParams",
                "ErrorCodeUserAlreadyExist",
                "ErrorCodeUserNotExist",
                "ErrorCodeInvalidPassword",
                "ErrorCodeServerBusy",
                "ErrorCodInvalidAuth",
                "ErrorCodeNeedLogin"
            ]
        },
        "controller._ResponseCommunity": {
            "type": "object",
            "properties": {
                "cname": {
                    "description": "社区名称",
                    "type": "string"
                },
                "code": {
                    "description": "业务响应状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controller.ResCode"
                        }
                    ]
                },
                "community_id": {
                    "description": "社区id",
                    "type": "integer"
                },
                "community_name": {
                    "description": "社区名称",
                    "type": "string"
                },
                "creattime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "introduction": {
                    "description": "描述信息",
                    "type": "string"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "controller._ResponsePost": {
            "type": "object",
            "required": [
                "community_id",
                "content",
                "title"
            ],
            "properties": {
                "author_id": {
                    "description": "作者id",
                    "type": "integer"
                },
                "author_name": {
                    "description": "作者名称",
                    "type": "string"
                },
                "code": {
                    "description": "业务响应状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controller.ResCode"
                        }
                    ]
                },
                "community_id": {
                    "description": "社区id",
                    "type": "integer"
                },
                "community_name": {
                    "description": "社区名称",
                    "type": "string"
                },
                "content": {
                    "description": "帖子内容",
                    "type": "string"
                },
                "creat_time": {
                    "description": "发帖时间",
                    "type": "string"
                },
                "creattime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "introduction": {
                    "description": "描述信息",
                    "type": "string"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                },
                "post_approve_num": {
                    "description": "帖子被点赞数",
                    "type": "integer"
                },
                "post_id": {
                    "description": "帖子id",
                    "type": "integer"
                },
                "status": {
                    "description": "帖子状态",
                    "type": "integer"
                },
                "title": {
                    "description": "帖子标题",
                    "type": "string"
                }
            }
        },
        "controller._ResponseUser": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/controller.ResCode"
                        }
                    ]
                },
                "date": {
                    "description": "返回数据"
                }
            }
        },
        "mod.ParamSignIn": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "mod.ParamSignUp": {
            "type": "object",
            "required": [
                "password",
                "re_password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "re_password": {
                    "description": "确认密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "mod.ParamsGetPostList": {
            "type": "object",
            "properties": {
                "community_id": {
                    "description": "可以为空",
                    "type": "integer"
                },
                "page": {
                    "description": "页码",
                    "type": "integer"
                },
                "size": {
                    "description": "每页数据量",
                    "type": "integer"
                },
                "type": {
                    "description": "排序依据",
                    "type": "string",
                    "example": "score"
                }
            }
        },
        "mod.Post": {
            "type": "object",
            "required": [
                "community_id",
                "content",
                "title"
            ],
            "properties": {
                "author_id": {
                    "description": "作者id",
                    "type": "integer"
                },
                "community_id": {
                    "description": "社区id",
                    "type": "integer"
                },
                "content": {
                    "description": "帖子内容",
                    "type": "string"
                },
                "creat_time": {
                    "description": "发帖时间",
                    "type": "string"
                },
                "post_id": {
                    "description": "帖子id",
                    "type": "integer"
                },
                "status": {
                    "description": "帖子状态",
                    "type": "integer"
                },
                "title": {
                    "description": "帖子标题",
                    "type": "string"
                }
            }
        },
        "mod.PostVoted": {
            "type": "object",
            "required": [
                "post_id"
            ],
            "properties": {
                "direction": {
                    "description": "投票类型，1赞同，-1不赞同，0取消投票",
                    "type": "integer",
                    "enum": [
                        1,
                        -1,
                        0
                    ]
                },
                "post_id": {
                    "description": "帖子id，必传",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Bluebell项目",
	Description:      "一个后端项目",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
