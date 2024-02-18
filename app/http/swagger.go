// Package http API.
// @title bbs 论坛
// @version 1.0
// @description hade测试
// @termsOfService https://github.com/swaggo/swag

// @contact.name zhangsan

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

package http

import (
	_ "bbs/app/http/swagger"
)
