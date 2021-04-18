package webcontainers

import (
	"bulma/web"
)

type IWebContainer interface {
	SetMarkdown(markdown string)

	web.IWeb
}
