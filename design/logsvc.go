package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("logsvc", func() {
	Description("Log aggregation service")

	Method("put", func() {
		Payload(func() {
			Field(1, "at", Int64, "Time of log entry")
			Field(2, "level", String, func() {
				Description("Level of log entry; one of debug, info, error")
				Default("info")
				Enum("info", "error", "debug")
			})
			Field(3, "fields", MapOf(String, String), "Log fields to tag message with")
			Field(4, "service", String, "Service providing the log message")
			Field(5, "message", String, "Actual log message")
			Required("at", "level", "service", "message")
		})
		GRPC(func() {})
	})
})
