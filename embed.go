// Copyright by AcmaTvirus
package foxdocker

import "embed"

//go:embed all:web/dist
var WebDist embed.FS

//go:embed all:templates
var TemplatesDist embed.FS
