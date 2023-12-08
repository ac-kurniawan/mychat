package http

import (
	"errors"

	"github.com/ac-kurniawan/mychat/core"
)

type HttpErrorModel struct {
	Code       error  `json:"code"`
	HttpStatus string `json:"httpStatus"`
	Message    string `json:"message"`
}

var HttpErrorList []HttpErrorModel = []HttpErrorModel{
	{
		Code:       core.ErrFailedGetRoomChat,
		HttpStatus: "500",
		Message:    "failed to get room chat record",
	},
	{
		Code:       core.ErrFailedGetParticipantGroup,
		HttpStatus: "500",
		Message:    "failed to get participant group record",
	},
	{
		Code:       core.ErrSessionStatusInvalid,
		HttpStatus: "400",
		Message:    "session status is not valid",
	},
}

func FindHttpError(err error) HttpErrorModel {
	var response HttpErrorModel
	for _, val := range HttpErrorList {
		if errors.Is(val.Code, err) {
			response = val
			return response
		}
	}
	return HttpErrorModel{
		Code:       errors.New("UNEXPECTED_ERROR"),
		HttpStatus: "500",
		Message:    "unexpected error",
	}
}
