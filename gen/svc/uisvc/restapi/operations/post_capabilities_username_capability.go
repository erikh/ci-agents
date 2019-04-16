// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinyci/ci-agents/clients/log"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/handlers"

	"github.com/google/uuid"
)

// PostCapabilitiesUsernameCapability swagger:route POST /capabilities/{username}/{capability} postCapabilitiesUsernameCapability
// Add a named capability// Add a named capability for a provided user ID. Requires the user have the 'modify:user' capability.
func PostCapabilitiesUsernameCapability(h *handlers.H, ctx *gin.Context, processingHandler handlers.HandlerFunc) *errors.Error {
	if h.RequestLogging {
		start := time.Now()
		u := uuid.New()

		content, jsonErr := json.Marshal(ctx.Params)
		if jsonErr != nil {
			h.Clients.Log.Error(errors.New(jsonErr).Wrap("encoding params for log message"))
		}

		logger := h.Clients.Log.WithRequest(ctx.Request).WithFields(log.FieldMap{
			"params":       string(content),
			"request_uuid": u.String(),
		})

		user, err := h.GetGithub(ctx)
		if err == nil {
			logger = logger.WithUser(user)
		}

		logger.Debug("incoming request")

		defer func() {
			logger.WithFields(log.FieldMap{
				"duration": time.Since(start).String(),
			}).Debug("request completed")
		}()
	}

	if err := PostCapabilitiesUsernameCapabilityValidateURLParams(h, ctx); err != nil {
		return errors.New(err)
	}

	if processingHandler == nil {
		return errors.New("'/capabilities/{username}/{capability}': no processor defined")
	}

	resp, code, err := processingHandler(h, ctx)
	return PostCapabilitiesUsernameCapabilityResponse(h, ctx, resp, code, err)
}
