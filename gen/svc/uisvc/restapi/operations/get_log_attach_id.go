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

// GetLogAttachID swagger:route GET /log/attach/{id} getLogAttachId
// Attach to a running log
// For a given ID, find the log and if it is running, attach to it and start receiving the latest content from it.
//
func GetLogAttachID(h *handlers.H, ctx *gin.Context, processingHandler handlers.HandlerFunc) *errors.Error {
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

	if err := GetLogAttachIDValidateURLParams(h, ctx); err != nil {
		return errors.New(err)
	}

	if processingHandler == nil {
		return errors.New("'/log/attach/{id}': no processor defined")
	}

	resp, code, err := processingHandler(h, ctx)
	return GetLogAttachIDResponse(h, ctx, resp, code, err)
}
