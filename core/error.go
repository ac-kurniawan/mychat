package core

import "errors"

var (
	ErrFailedGetRoomChat         error = errors.New("ErrFailedGetRoomChat")
	ErrFailedGetParticipantGroup error = errors.New("ErrFailedGetParticipantGroup")
	ErrSessionStatusInvalid      error = errors.New("ErrSessionStatusInvalid")
)
