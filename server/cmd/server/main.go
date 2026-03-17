package main

import (
	"github.com/InkSha/Resummer/internal/api"
	"github.com/InkSha/Resummer/internal/initialize"
)

func bootstrap() {
	initialize.Initialize()
	api.Listen("api", 8080)
}

//	@title			Resummer API
//	@version		0.0.1
//	@description	This is a Resummer API server.
//	@termsOfService	http://swagger.io/terms/

//	@License					MIT
//	@BasePath					/
//	@host						localhost:8080
//	@scheme						http https
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@description				Type "Bearer" followed by a space and JWT token.
//	@description				```
//	@description				eg:
//	@description				Bearer <JWT Token>
//	@description				```
//	@in							header
//	@name						Authorization
//	@type						apiKey

func main() {
	bootstrap()
}
