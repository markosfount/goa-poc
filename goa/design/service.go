package design

import (
	. "goa.design/goa/v3/dsl"
)

var NotFoundErr = Type("NotFoundError", func() {
	Attribute("message", String)
	Required("message")
})

var _ = API("poc", func() {
	Title("poc Service")
	Description("foo bar")

	Server("poc", func() {
		Description("poc server")
		Services("poc")

		Host("development", func() {
			Description("Development hosts.")
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("poc", func() {
	Description("poc")
	Error("not_found", NotFoundErr)

	Method("get", func() {
		Payload(func() {
			Attribute("id", Int, "ID of the item")
			Required("id")
		})

		Result(HttpConfig)

		HTTP(func() {
			GET("/todos/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

})
