package uisvc

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/tinyci/ci-agents/ci-gen/openapi/services/uisvc"
	"github.com/tinyci/ci-agents/utils"
)

// GetRunsCount returns a count of the queue items by asking the datasvc for it.
func (h *H) GetRunsCount(ctx echo.Context, params uisvc.GetRunsCountParams) error {
	count, err := h.clients.Data.RunCount(ctx.Request().Context(), stringDeref(params.Repository), stringDeref(params.Sha))
	if err != nil {
		return err
	}

	return ctx.JSON(200, count)
}

// GetRuns lists all the runs that were requested by the page/perPage parameters.
func (h *H) GetRuns(ctx echo.Context, params uisvc.GetRunsParams) error {
	page, perPage, err := utils.ScopePaginationInt(params.Page, params.PerPage)
	if err != nil {
		return err
	}

	list, err := h.clients.Data.ListRuns(ctx.Request().Context(), stringDeref(params.Repository), stringDeref(params.Sha), int64(page), int64(perPage))
	if err != nil {
		return err
	}

	r, err := h.convertRuns(ctx, list)
	if err != nil {
		return err
	}

	return ctx.JSON(200, r)
}

// GetRunRunId retrieves a run by id.
func (h *H) GetRunRunId(ctx echo.Context, runID int64) error {
	run, err := h.clients.Data.GetRunUI(ctx.Request().Context(), runID)
	if err != nil {
		return err
	}

	r, err := h.C.FromProto(ctx.Request().Context(), run)
	if err != nil {
		return err
	}

	return ctx.JSON(200, r)
}

// PostCancelRunId cancels a run by id.
func (h *H) PostCancelRunId(ctx echo.Context, runID int64) error {
	if err := h.clients.Data.SetCancel(context.Background(), runID); err != nil {
		return err
	}

	return ctx.NoContent(200)
}
