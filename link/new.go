package link

import "github.com/hashicorp/go-plugin"

func New() plugin.Plugin {
	return &link{}
}
