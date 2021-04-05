// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/tinyci/ci-agents/clients/log"
	"github.com/tinyci/ci-agents/handlers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetLoginUpgrade swagger:route GET /login/upgrade getLoginUpgrade
// Log into the system with upgraded permissions
// This upgrades the permissions of the user (which requires confirmation from the OAuthing site) to allow repository access, so that additional permission to manipulate repositories and scan additional ones is available.
//
func GetLoginUpgrade(h *handlers.H, ctx *gin.Context, processingHandler handlers.HandlerFunc) error {
	if h.RequestLogging {
		start := time.Now()
		u := uuid.New()

		content, jsonErr := json.Marshal(ctx.Params)
		if jsonErr != nil {
			h.Clients.Log.Error(ctx.Request.Context(), fmt.Errorf("encoding params for log message: %w", jsonErr))
		}

		logger := h.Clients.Log.WithRequest(ctx.Request).WithFields(log.FieldMap{
			"params":       string(content),
			"request_uuid": u.String(),
		})

		user, err := h.GetGithub(ctx)
		if err == nil {
			logger = logger.WithUser(user)
		}

		logger.Debug(ctx.Request.Context(), "incoming request")

		defer func() {
			logger.WithFields(log.FieldMap{
				"duration": time.Since(start).String(),
			}).Debug(ctx.Request.Context(), "request completed")
		}()
	}

	if err := GetLoginUpgradeValidateURLParams(h, ctx); err != nil {
		return err
	}

	if processingHandler == nil {
		return errors.New("'/login/upgrade': no processor defined")
	}

	processingContext, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	resp, code, err := processingHandler(processingContext, h, ctx)
	return GetLoginUpgradeResponse(h, ctx, resp, code, err)
}
