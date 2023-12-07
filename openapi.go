package openapi

import (
	"github.com/aCt0802/openapi-gen/info"
	"github.com/aCt0802/openapi-gen/paths"
	"github.com/aCt0802/openapi-gen/servers"
	"github.com/aCt0802/openapi-gen/tags"
)

type OpenAPI struct {
	openapi string
	Info    info.Info
	servers servers.Servers
	Tags    tags.Tags
	Paths   paths.Paths
}

func NewOpenAPI() *OpenAPI {
	return &OpenAPI{openapi: "3.0.3"}
}
