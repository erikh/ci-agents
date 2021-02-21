package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("assetsvc", func() {
	Description("Asset routing and aggregation service")

	Method("putLog", func() {
		StreamingPayload(func() {
			Field(1, "id", Int64, "Run identitfier")
			Field(2, "chunk", Bytes)
			Required("id")
		})
		GRPC(func() {})
	})

	Method("getLog", func() {
		Payload(func() {
			Field(1, "id", Int64, "Run identitfier")
			Required("id")
		})
		StreamingResult(func() {
			Field(1, "chunk", Bytes)
			Field(2, "errors", ArrayOf(String))
		})
		GRPC(func() {})
	})
})
