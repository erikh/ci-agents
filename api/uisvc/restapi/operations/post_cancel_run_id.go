// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"time"

	"github.com/tinyci/ci-agents/clients/log"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/handlers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PostCancelRunID swagger:route POST /cancel/{run_id} postCancelRunId
// Cancel by Run ID
// Cancel the run by ID; this will actually trickle back and cancel the whole task, since it can no longer succeed in any way.
// Please keep in mind to stop runs, runners must implement a cancel poller.
//
func PostCancelRunID(h *handlers.H, ctx *gin.Context, processingHandler handlers.HandlerFunc) *errors.Error {
	if h.RequestLogging {
		start := time.Now()
		u := uuid.New()

		content, jsonErr := json.Marshal(ctx.Params)
		if jsonErr != nil {
			h.Clients.Log.Error(ctx.Request.Context(), errors.New(jsonErr).Wrap("encoding params for log message"))
		}

		logger := h.Clients.Log.WithRequest(ctx.Request).WithFields(log.Fields{
			"params":       string(content),
			"request_uuid": u.String(),
		})

		user, err := h.GetGithub(ctx)
		if err == nil {
			logger = logger.WithUser(user)
		}

		logger.Debug(ctx.Request.Context(), "incoming request")

		defer func() {
			logger.WithFields(log.Fields{
				"duration": time.Since(start).String(),
			}).Debug(ctx.Request.Context(), "request completed")
		}()
	}

	if err := PostCancelRunIDValidateURLParams(h, ctx); err != nil {
		return errors.New(err)
	}

	if processingHandler == nil {
		return errors.New("'/cancel/{run_id}': no processor defined")
	}

	processingContext, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	resp, code, err := processingHandler(processingContext, h, ctx)
	return PostCancelRunIDResponse(h, ctx, resp, code, err)
}
