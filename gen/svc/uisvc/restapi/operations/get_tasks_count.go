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

// GetTasksCount swagger:route GET /tasks/count getTasksCount
// Count the Tasks
// Perform a full count of tasks that meet the filter criteria (which can be no filter) and return it as integer.
//
func GetTasksCount(h *handlers.H, ctx *gin.Context, processingHandler handlers.HandlerFunc) *errors.Error {
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

	if err := GetTasksCountValidateURLParams(h, ctx); err != nil {
		return errors.New(err)
	}

	if processingHandler == nil {
		return errors.New("'/tasks/count': no processor defined")
	}

	resp, code, err := processingHandler(h, ctx)
	return GetTasksCountResponse(h, ctx, resp, code, err)
}
