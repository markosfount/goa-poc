package design

import (
	. "goa.design/goa/v3/dsl" //nolint:revive
	"goa_poc/pkg"
)

var HttpConfig = Type("HttpConfig", func() {
	ConvertTo(pkg.HttpConfig{})
	Attribute("host", String, "The host")
	Attribute("port", Int, "The port number")
})
