package http

import (
	"errors"

	"github.com/ac-kurniawan/mychat/core"
	"github.com/ac-kurniawan/mychat/interface/http/dto"
	"github.com/ac-kurniawan/mychat/library"
	"github.com/labstack/echo/v4"
)

type HttpHandler struct {
	Service core.IMychatService
	Trace   library.AppTrace
}

// GetRoomChatBySessionId
// @Summary	get room chat by session ID
// @Tags	roomChat
// @Accept		json
// @Produce	json
// @Param		session_id	query		string	true	"session id"
// @Param		session_status	query		string	false	"filter session status"
// @Success	200		{object}	dto.RoomChatDto
// @Failure 500 {object} HttpErrorModel
// @Failure 400 {object} HttpErrorModel
// @Router		/room_chat [get]
func (h HttpHandler) GetRoomChatBySessionId(c echo.Context) error {
	ctx, span := h.Trace.StartTraceServer(c.Request().Context(), "HTTP SERVER - GetRoomChatBySessionId")
	defer h.Trace.EndTrace(span)
	var query dto.GetRoomChatBySessionIdDto
	err := c.Bind(&query)
	if err != nil {
		return c.JSON(400, HttpErrorModel{
			Code:       errors.New("VALIDATION_ERROR"),
			HttpStatus: 400,
			Message:    "query not valid",
		})
	}

	result, err := h.Service.GetRoomChatBySessionId(ctx, query.SessionId, query.SessionStatus)
	if err != nil {
		errMsg := FindHttpError(err)
		return c.JSON(errMsg.HttpStatus, errMsg)
	}

	var response dto.RoomChatDto
	response.FromModel(*result)

	return c.JSON(200, response)
}
