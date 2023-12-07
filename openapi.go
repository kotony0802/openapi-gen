package openapi

import (
	"github.com/aCt0802/openapi-gen/info"
	"github.com/aCt0802/openapi-gen/paths"
	"github.com/aCt0802/openapi-gen/servers"
	"github.com/aCt0802/openapi-gen/tags"
)

type OpenAPI struct {
	Info    info.Info
	servers servers.Servers
	Tags    tags.Tags
	Paths   paths.Paths
}
