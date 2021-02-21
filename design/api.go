package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("tinyci", func() {
	Title("tinyCI")
	Description("Continuous Integration Service")
	Server("demo", func() {
		Services("logsvc", "assetsvc")

		Host("localhost", func() {
			URI("http://localhost:3000/uisvc")
		})
	})
})
