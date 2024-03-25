package option

import (
	"github.com/coorify/backend"
	"github.com/coorify/backend/option"
	"github.com/coorify/quickstart/plugin"
	"github.com/coorify/quickstart/router"
)

type Option struct {
	option.Option `anonymous:"true" yaml:",inline"`
}

func (o *Option) Plugin(s *backend.Server) error {
	plugin.Setup(s)
	return nil
}

func (o *Option) Router(s *backend.Server) error {
	router.Setup(s)
	return nil
}
