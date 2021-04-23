package webcontainers

import (
	"bulma/web"
)

type IWebContainer interface {
	SetContent(content []string)

	web.IWeb
}
